package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rpgmypet/internal/databases"
	"rpgmypet/internal/handlers"

	"github.com/gorilla/mux"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type HttpServer struct {
	Environment string
	Port        string
	DBConfig    *DBConfig
	Router      *mux.Router
}

func NewHttpServer(ctx context.Context, environment, port string, dbCfg *DBConfig) *HttpServer {
	return &HttpServer{
		Environment: environment,
		Port:        port,
		DBConfig:    dbCfg,
		Router:      mux.NewRouter(),
	}
}

func (srv *HttpServer) SetEndpoints() {
	// swagger:operation POST /creamascota pet CreatePet
	// Añade una nueva 'mascota' al sistema
	// Agrega una nueva entidad 'pet' al sistema
	// ---
	// produces:
	// - application/json
	// Parameters:
	// - name: name
	//   in: body
	//   description: Nombre de la nueva mascota
	//   required: true
	//   schema:
	//     type: string
	// - name: specie
	//   in: body
	//   description: Denominación de la especie de la nueva mascota
	//   required: true
	//   schema:
	//     type: string
	// - name: sex
	//   in: body
	//   description: Sexo de la nueva mascota (Masculino o femenino)
	//   required: true
	//   schema:
	//     type: string
	// - name: birthdate
	//   in: body
	//   description: Fecha de nacimiento en formato 'yyyy-mm-ddThh:mm:ssZ'
	//   required: true
	//   schema:
	//     type: string
	// Responses:
	//   '200':
	//     description: Éxito. Ok, sin problema
	//   '400':
	//     description: Error. Petición mal realizada
	//   '500':
	//     description: Error. Interno del servidor
	srv.Router.HandleFunc("/creamascota", handlers.CreatePet()).Methods(http.MethodPost)

	// swagger:operation GET /lismascotas pet ListPets
	// Listado de las 'mascotas' registradas en el sistema
	// Proporciona un listado de las entidades 'Pet' en el sistema.
	// ---
	// produces:
	// - application/json
	// Responses:
	//   '200':
	//     description: Éxito. Ok, sin problema
	//   '400':
	//     description: Error. Petición mal realizada
	//   '500':
	//     description: Error. Interno del servidor
	srv.Router.HandleFunc("/lismascotas", handlers.ListPets()).Methods(http.MethodGet)

	// swagger:operation GET /kpidemascotas pet ReportPets
	// Facilita la especie mas abundante entre las entidades 'Pet' registradas en el sistema.
	// Proporciona la especie con mas individuos registrados en el sistema
	// ---
	// produces:
	// - application/json
	// Responses:
	//   '200':
	//     description: Éxito. Ok, sin problema
	//   '400':
	//     description: Error. Petición mal realizada
	//   '500':
	//     description: Error. Interno del servidor
	srv.Router.HandleFunc("/kpidemascotas", handlers.ReportPets()).Methods(http.MethodGet)

	// swagger:operation GET /kpidemascotas/{specie} pet ReportPetsSpecies
	// Proporciona datos de una especia de 'Pet' dada específica
	// Facilita la edad media y la desviación estándar de las edades de las 'Pet' registradas
	// ---
	// produces:
	// - application/json
	// Parameters:
	// - name: specie
	//   in: path
	//   description: Denominación de la especie de la nueva mascota
	//   required: true
	//   type: string
	// Responses:
	//   '200':
	//     description: Éxito. Ok, sin problema
	//   '400':
	//     description: Error. Petición mal realizada
	//   '500':
	//     description: Error. Interno del servidor
	srv.Router.HandleFunc("/kpidemascotas/{specie}", handlers.ReportPets()).Methods(http.MethodGet)

	srv.Router.PathPrefix("/doc/").Handler(http.StripPrefix("/doc/", http.FileServer(http.Dir("./internal/swaggerui"))))
}

//* Método para configurar, iniciar y comenzar a escuchar peticiones HTTP.
func (srv *HttpServer) Start() {
	dbr, err := databases.NewPostgresImplementation(
		srv.Environment,
		srv.DBConfig.User,
		srv.DBConfig.Password,
		srv.DBConfig.Host,
		srv.DBConfig.Port,
		srv.DBConfig.Name,
	)
	if err != nil {
		log.Fatalf("Database connection failed: '%v'", err)
	}

	databases.SetDatabaseRepository(dbr)

	srv.SetEndpoints()

	log.Printf("Starting server on port: %s\n", srv.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), srv.Router); err != nil {
		log.Fatalf("Failed 'ListenAndServe': %v", err)
	}
}
