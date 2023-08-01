# GO REST API

## Ejecucion
Descarga el proyecto y dentro de la carpeta de proyecto descargado ejecuta el comando 
```bash
go run main.go
```
Se levantara un servidor en el puerto 8080

## Uso
Aqui hay un ejemplo de las peticiones que se pueden mandar:

Para obtener todos los registros: GET localhost:8080/tasks
Para obtener un solo registro: GET localhost:8080/task/{id} (cambiar id por un valor tipo integer)
Para eliminar un registro: DELETE localhost:8080/task/{id} (cambiar id por un valor tipo integer)
Para crear un nuevo registro: POST localhost:8080/task (en el body se tiene que enviar un json con la siguiente informacion)
```json
{
	"name": "Task 3",
	"desc": "Take a shower",
	"completed": true
}
```
