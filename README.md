### Para comenzar a testear, ejecutar:
> `docker-compose build && docker-compose up -d`

> `go run ./cmd`

### Los endpoints son:
- '/creamascota', que solicita un objeto json tal que así:
<pre>{
	"name": "nombre18",
	"specie": "especie18",
	"sex": "sexo18",
	"birthdate": "2022-09-21T22:10:11Z"
}</pre>

- '/lismascotas', que solicita un objeto json tal que así:
<pre>{
	"pageInfo":{
		"orderBy": "name asc",
		"specie": "",
		"pageSize": "",
		"pageToken": 0,
		"totalPages": "",
		"totalItems": "",
    }
}</pre>

- '/kpidemascotas/{specie}', que solicita, o no, un queryParam llamado "specie" que, en caso de concurrencia, predomina con respecto al siguiente endpoint.

- '/kpidemascotas', que solicita, o no, un objeto json tal que así:
<pre>{
	"specie": ""
}</pre>

- '/doc', que muestra la documentación precompilada con go-swagger

### Para finalizar:
- pulsar 'Ctrl+C' en la consola para detener la ejecución de Golang
- y, para detener y limpiar virtualización Docker, ejecutar:
> `docker-compose down --rmi local`
