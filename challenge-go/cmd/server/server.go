package server

import (
	"challenge-go/cmd/tools/env"
	"challenge-go/pkg/middleware"
	"challenge-go/pkg/router"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {

	var server = Server{}
	app := server.initializeServer()
	server.App = app

	// Middlewares
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.HealthCheck())

	// Routes
	router.MainRoutes(app)

	if env.IsProd() {
		server.startServerWithGracefulShutdown()
	} else {
		server.startServer()
	}

	return &server
}

func (Server) initializeServer() *fiber.App {
	var timeout = env.GetEnv(env.ServerTimeout)
	var name = env.GetEnv(env.ServerName)

	readTimeoutSecondsCount, _ := strconv.Atoi(timeout)

	// Return Fiber configuration.
	config := fiber.Config{
		ServerHeader: name,
		AppName:      name,
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler: DefaultErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		//DisableStartupMessage: true,
	}
	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(loggerRequest())
	app.Use(NewRelicConfig())

	// @Deprecated
	//app.Use(routers.SwaggerHandler())

	return app
}

func loggerRequest() fiber.Handler {
	return logger.New(logger.Config{
		TimeFormat: "2006/01/02 15:04:05.999999",
		TimeZone:   "America/Lima",
		//Format:       "[uuid: ${locals:requestid} path: ${path} method: ${method} status: ${status} response: ${resBody}]\n",
		Format:       "[uuid: ${locals:requestid} method: ${method} status: ${status} time: ${time} latency: ${latency} ip: ${ip} path: ${path}]\n",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
		Done: func(c *fiber.Ctx, logString []byte) {
			if strings.Contains(c.Path(), "/api") {
				//if strings.EqualFold(c.Method(), fiber.MethodGet) {
				//	fmt.Printf("[uuid: %s method: %s full_url: %s ]\n", c.Locals("requestid"), c.Path(), c.Method(), c.OriginalURL())
				//}
				//if strings.EqualFold(c.Method(), fiber.MethodPost) || strings.EqualFold(c.Method(), fiber.MethodPut) || strings.EqualFold(c.Method(), fiber.MethodDelete) {
				//	fmt.Printf("[uuid: %s path: %s method: %s full_url: %s body: %s ]\n", c.Locals("requestid"), c.Path(), c.Method(), c.OriginalURL(), c.Request().Body())
				//}
				fmt.Printf("[uuid: %s method: %s status: %d path: %s response: %s ]\n", c.Locals("requestid"), c.Method(), c.Response().StatusCode(), c.Path(), c.Response().Body())
			}
		},
	})
}

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	var message string
	var e *fiber.Error
	var code int
	if env.IsLocal() {
		if errors.As(err, &e) {
			code = e.Code
		}
		message = err.Error()
	} else {
		message = "Datos insuficientes para el servicio"
	}

	if code == fiber.StatusNotFound {
		message = "No se encontró la ruta solicitada"
	}

	if code == fiber.StatusInternalServerError {
		message = "Error interno del servidor"
	}

	return c.Status(code).SendString(message)
}

func (s Server) startServerWithGracefulShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.App.Shutdown(); err != nil {
			fmt.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := s.App.Listen(GetPort()); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
		//fmt.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func (s Server) startServer() {
	if err := s.App.Listen(GetPort()); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
		//fmt.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func GetPort() string {
	var port = env.GetEnv(env.ServerPort)
    return fmt.Sprintf(":%s", port)
}
