# Coding Challenge Node
![challenge.png](./../_images/challenge.png)
 
 _**Nombre:** Yahyr Paredes Arteaga_ @yahyrparedes
 

## Requerimientos:
1. [x] Node 22 o superior (Recomendado)
2. [x] npm 6 o superior
3. [x] Docker 20 o superior (Opcional)
4. [x] Express.js 4.17.1

## Comandos:
1. [x] `npm start` para iniciar la aplicación
2. [x] `npm test` para ejecutar las pruebas unitarias
3. [x] `npm run dev` para iniciar la aplicación en modo desarrollo
4. [x] `npm run build` para compilar la aplicación
5. [x] `npm run lint` para ejecutar el linter

## Instalación:
1. Clonar el repositorio
```bash
git clone 
```
2. Ingresar al directorio
```bash
cd challenge-node
```
3. Instalar las dependencias
```bash  
npm install
```
4. Iniciar la aplicación
```bash
npm start
```
5. Probar la siguiente URL
```bash
http://localhost:3002/api/v1/challenge
```

## Docker:
1. Construir la imagen
```bash
docker build -t challenge-node .
```
2. Iniciar el contenedor
```bash
docker run -p 3002:3002 challenge-node
```
3. Probar la siguiente URL
```bash
http://localhost:3002/api/v1/challenge
```

## Endpoints: 
1. [x] POST /api/v1/challenge

## Ejemplo:
```bash
curl --location --request POST 'http://localhost:3002/api/v1/challenge' \
--header 'Content-Type: application/json' \
--data-raw '{ 
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
    ]     
}'
```

## Respuesta:
```json
{
    "status": 200,
    "results": {
        "min": -0.345032,
        "max": 7.437357,
        "avg": 2.1642484166666667,
        "sum": 16.530869,
        "isDiagonalQ": true,
        "isDiagonalR": true
    }
}
```
