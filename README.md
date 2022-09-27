### Requisitos (tener instalado):
- Docker y docker-compose
- Go

### Para comenzar a realizar pruebas en entorno de desarrollo local, ejecutar:
> `make serve`

### Para finalizar las pruebas en entorno de desarrollo local, ejecutar:
> `make destroy`

### Para comenzar a realizar pruebas en entorno de desarrollo local, visitar (dicha URL la denominaremos '~'):
> `http://localhost:8080`

### Para comenzar a realizar pruebas en entorno de producción, únicamente visitar (dicha URL la denominaremos '~'):
> `https://rpgmypet-server-mh2nagbpxq-no.a.run.app`

### Los endpoints son:
- '~/creamascota', que solicita un objeto json tal que así:
<pre>{
	"name": "nombre18".
	"specie": "especie18",
	"sex": "sexo18",
	"birthdate": "2022-09-21T22:10:11Z" (Necesariamente debe ser éste formato: 'yyyy-mm-ddThh:mm:ssZ')
}</pre>

- '~/lismascotas'

- '~/kpidemascotas/{specie}', que solicita un queryParam llamado "specie" que, en caso de concurrencia,
predomina con respecto al siguiente endpoint.

- '~/kpidemascotas'

- '~/doc/', que muestra la documentación precompilada con go-swagger
