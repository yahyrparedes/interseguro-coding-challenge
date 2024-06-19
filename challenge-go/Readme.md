# Coding Challenge Node
![challenge.png](./../_images/challenge.png)

_**Nombre:** Yahyr Paredes Arteaga_ @yahyrparedes
 
## Requerimientos:
1. [x]  Go 1.22
2. [x]  Docker
3. [x]  swag (go get -u github.com/swaggo/swag/cmd/swag)
4. [x]  Fiber
5. [x]  New Relic

## Instalación
1. Clonar el repositorio
```bash
git clone
```
2. Instalar las dependencias
```bash
go mod download
```
3. Crear el archivo .env
```bash
cp .env.example .env
```
4. Ejecutar el proyecto
```bash
swag init -g cmd/main.go --output docs
go run main.go
```
5. Acceder a la documentación
```bash
http://localhost:3000/swagger/index.html
```


## Docker:
1. Crear la imagen
```bash
docker build -t challenge-go .
```
2. Ejecutar el contenedor
```bash
docker run -p 8080:8080 challenge-go
```
3. Acceder a la documentación
```bash 
http://localhost:8080/swagger/index.html
```

## Endpoints
1. [x]  GET /api/v1/health
2. [x]  POST /api/v1/challenge 

## Ejemplo de uso
```bash
curl --location --request POST 'http://localhost:8080/api/v1/challenge' \
--header 'Content-Type: application/json' \
--data-raw '{
  "matrix" :  [
        [1,2],
        [3,4],
        [5,6]
    ]  
}'
```

## Respuesta
```json
{
  "message": "Success",
  "status": 200,
  "data": {
    "q": [
      [
        0.16903,
        0.897085
      ],
      [
        0.507092,
        0.276026
      ],
      [
        0.845154,
        -0.345032
      ]
    ],
    "r": [
      [
        5.916079,
        7.437357
      ],
      [
        0,
        0.828078
      ]
    ],
    "qx": [
      [
        0.16903,
        0.897085
      ],
      [
        0.507092,
        0.276026
      ],
      [
        0.845154,
        -0.345032
      ]
    ],
    "rx": [
      [
        5.916079,
        7.437357
      ],
      [
        0,
        0.828078
      ]
    ]
  }
}
```


## Estructura del proyecto
```bash

├── cmd
│   └── main.go 
│   ├── server
├── ├── ├── newrelic.go
├── ├── ├── server.go
├── ├── tools
├── ├── ├── tools.go
├── ├── ├── env
├── ├── ├── ├── env.go
├── ├── ├── ├── dynamic.go
├── ├── ├── log
├── ├── ├── ├── log.go
├── docs
│   ├── docs.go
│   └── swagger.json
├── pkg
│   ├── middleware
│   │   └── cors.go
│   ├── router
│   │   └── router.go
│   ├   ├── swagger.go
│   ├── tools
│   │   └── tools.go
├── api
│   ├── v1
│   │   ├── challenge
│   │   │   ├── controller.go
│   │   │   ├── dto.go
│   │   │   ├── router.go 
│   ├── ├── commons
│   │   │   ├── commons.go
├── test
│   ├── test.go

```


