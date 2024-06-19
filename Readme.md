# Coding Challenge - 2024

![challenge.png](./_images/challenge.png)

### DIVISIÓN TI 

 _**Nombre:** Yahyr Paredes Arteaga_ @yahyrparedes

## Desafío Técnico - Descripción


### Consideraciones técnicas:
1. [x] Utilizar el lenguaje de programación Go (Golang) para una API y Node.js para la otra API.
2. [x] Implementar la solución utilizando los frameworks Fiber para la API en Go y Express.js para la API en Node.js.**
3. [x] Documentar el código de manera clara y concisa, siguiendo las mejores prácticas de codificación.
4. [x] Utilizar Docker para contenerizar las aplicaciones y facilitar su despliegue en diferentes entornos.**
5. [x] Implementar la comunicación entre las dos API utilizando un mecanismo como HTTP.
   - Utilizar servicios en la nube para la implementación y el despliegue de las aplicaciones.

###  Arquitectura de la solución:

1. [x] API en Go: Esta API recibirá la matriz original como entrada, realizará la rotación de la matriz y luego enviará los datos resultantes a la segunda API en Node.js.
2. [x] API en Node.js: Esta API recibirá los datos de la matriz rotada de la API en Go, calculará estadísticas sobre los datos y devolverá estas estadísticas como resultado.


### Funcionalidad requerida:

1. [x] Crear dos API RESTful:
   - Una API en Go que reciba como entrada un array de arrays de números que represente una matriz rectangular y devuelva la factorización QR de dicha matriz.
   - Otra API en Node.js que reciba el resultado de las matrices devueltas por la primera API y realice una operación adicional sobre los datos. (*) Detalle en la sección operaciones adicionales
2. [x] Implementar la lógica para realizar la rotación de la matriz y la operación adicional de manera eficiente y correcta en cada API.

### Funcionalidad opcional:
1. [x] Implementar un frontend que consuma ambas APIs y muestre los resultados de la rotación de la matriz y las estadísticas adicionales.
2. [x] Aplicar un nivel de seguridad utilizando JWT para proteger las consultas a las APIs.
3. [x] Implementar pruebas unitarias y de integración para garantizar la calidad del código en ambas API.

### Operación adicional:

1. [x] La segunda API calculará lo siguientes sobre los datos de las matrices devueltas:
   - Valor máximo: El valor máximo encontrado en las matrices.
   - Valor mínimo: El valor mínimo encontrado en las matrices.
   - Promedio: El promedio de todos los valores de las matrices.
   - Suma total: La suma total de todos los valores de las matrices.
   - Matriz diagonal: Verificar si alguna matriz es diagonal.

### Consideraciones:
    
1. [x] No hay un estándar específico para los nombres de los objetos creados, pero se espera coherencia en su estructura y documentación.
2. [x] En caso de dudas en el enunciado, se espera que el candidato tome decisiones informadas y las sustente durante la entrevista.
3. [x] Se valorará la eficiencia y la elegancia de la solución implementada, así como la capacidad del candidato para comunicar y defender sus decisiones técnicas.



 
![logo.png](./_images/logo.png) 

## Gracias por la oportunidad de participar en este desafío técnico. 🚀
