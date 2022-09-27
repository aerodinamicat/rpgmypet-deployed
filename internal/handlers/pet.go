package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"rpgmypet/internal/databases"
	"rpgmypet/internal/models"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

const (
	NOT_FOUND_SPECIE = "Specie not found"
)

type createRequest struct {
	Name      string    `json:"name"`
	Specie    string    `json:"specie"`
	Sex       string    `json:"sex"`
	Birthdate time.Time `json:"birthdate"`
}

// swagger:response createPetResponse
type createResponse struct {
	// in: body
	Pet models.Pet `json:"pet"`
}

// swagger:response listPetResponse
// description: DTO para contestar a la petición de 'ListPetRequest'
type listResponse struct {
	// in: body
	Pets []*models.Pet `json:"pets"`
}

// swagger:response reportPetResponse
type reportResponse struct {
	// in: body
	Specie string `json:"specie"`

	AverageAge        string `json:"averageAge,omitempty"`
	StandardDeviation string `json:"standardAgeDeviation,omitempty"`
}

//* Función que será ejecutada en el 'endpoint' -> POST '/creamascota'
func CreatePet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var decodedRequest = new(createRequest)
		if err := json.NewDecoder(request.Body).Decode(&decodedRequest); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		newId, err := ksuid.NewRandom()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		pet := models.Pet{
			Name:      decodedRequest.Name,
			Specie:    decodedRequest.Specie,
			Sex:       decodedRequest.Sex,
			Birthdate: decodedRequest.Birthdate,
			Id:        newId.String(),
		}

		if err := databases.InsertPet(request.Context(), &pet); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		notEncodedResponse := createResponse{
			Pet: pet,
		}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(notEncodedResponse)
	}
}

//* Función que será ejecutada en el 'endpoint' -> GET '/lismascotas'
func ListPets() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pets, err := databases.ListPets(request.Context(), "")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		notEncodedResponse := listResponse{
			Pets: pets,
		}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(notEncodedResponse)
	}
}

//* Función que será ejecutada en el 'endpoint' -> GET '/kpidemascotas'
func ReportPets() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryParams := mux.Vars(request)
		if filterBySpecie := queryParams["specie"]; filterBySpecie != "" {

			pets, err := databases.ListPets(request.Context(), filterBySpecie)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			if len(pets) == 0 {
				http.Error(writer, NOT_FOUND_SPECIE, http.StatusNotFound)
				return
			}

			avgAge, stDev := getAVGAgeAndSTDeviation(pets)
			notEncodedResponse := reportResponse{
				Specie:            filterBySpecie,
				AverageAge:        fmt.Sprintf("%f days", avgAge),
				StandardDeviation: fmt.Sprintf("%f days", stDev),
			}

			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(notEncodedResponse)
			return
		}

		mostCommonSpecie, err := databases.MostCommonPetSpecie(request.Context())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		notEncodedResponse := reportResponse{
			Specie: mostCommonSpecie,
		}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(notEncodedResponse)
	}
}

/** Función, de ámbito privado, que realiza los cálculos requeridos en el challenge sobre las edades
*   de las entidades 'Pet' del sistema.
 */
func getAVGAgeAndSTDeviation(items []*models.Pet) (float64, float64) {
	var sumAges float64
	for _, item := range items {
		sumAges += item.GetAgeInDays()
	}
	avgAge := sumAges / float64(len(items))

	var stDev float64
	for _, item := range items {
		stDev += math.Pow(item.GetAgeInDays()-avgAge, 2)
	}

	stDev = math.Sqrt(stDev / float64(len(items)))

	return avgAge, stDev
}
