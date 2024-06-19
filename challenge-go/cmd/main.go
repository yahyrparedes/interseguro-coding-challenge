package main

import (
	"challenge-go/cmd/server"
	"challenge-go/cmd/tools"
	"challenge-go/cmd/tools/env"
	_ "challenge-go/docs"
)

// @title Challenge API
// @description Challenge API
// @schemes http https
// @BasePath /

// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email yahyrparedesarteaga@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	// Environment
	env.NewEnv()
	// Tools
	tools.ConfigureTimeZone()
	// Server
	server.NewServer()

}
