### Para comenzar a testear, visitar (dicha URL la denominaremos '~'):
> `https://rpgmypet-server-mh2nagbpxq-no.a.run.app`

### Los endpoints son:
- '~/creamascota', que solicita un objeto json tal que así:
<pre>{
	"name": "nombre18".
	"specie": "especie18",
	"sex": "sexo18",
	"birthdate": "2022-09-21T22:10:11Z"
}</pre>

- '~/lismascotas', que solicita queryParams:
	- 'orderBy': en formato SQL, dicta en qué orden (ascendente o descendente) y qué campo, o columna,
	deben ordenarse los resultados.
	- 'filterBySpecie: barema resultados por dicho campo'.
	- 'pageSize': ajusta la cantidad de resultados mostrados en cada muestreo.
	- 'pageToken': visualiza que grupo de resultados, según 'pageSize' debe mostrarse. Corresponde al
	número de página a mostrar.
	- 'totalItems': corresponde a la cantidad de resultados encontrados en la búsqueda.
	- 'totalPages': corresponde a la cantidad de muestreos máximos para visualizar todos los resultados.

- '~/kpidemascotas/{specie}', que solicita un queryParam llamado "specie" que, en caso de concurrencia,
predomina con respecto al siguiente endpoint.

- '~/kpidemascotas'

- '~/doc', que muestra la documentación precompilada con go-swagger
