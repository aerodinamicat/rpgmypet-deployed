package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"rpgmypet/internal/databases"
	"rpgmypet/internal/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

const (
	DEFAULT_PAGE_SIZE = "10"
	DEFAULT_ORDER_BY  = "name asc"
)

type createRequest struct {
	Name      string    `json:"name"`
	Specie    string    `json:"specie"`
	Sex       string    `json:"sex"`
	Birthdate time.Time `json:"birthdate"`
}

// swagger:response createPetResponse
type createResponse struct {
	// in:body
	Pet models.Pet `json:"pet"`
}

// swagger:response listPetResponse
// description: DTO para contestar a la petición de 'ListPetRequest'
type listResponse struct {
	// in:body
	PageInfo *models.Pagination `json:"pageInfo"`
	Pets     []*models.Pet      `json:"pets"`
}

// swagger:response reportPetResponse
type reportResponse struct {
	// in:body
	Specie string `json:"specie"`

	AverageAge        string `json:"averageAge,omitempty"`
	StandardDeviation string `json:"standardAgeDeviation,omitempty"`
}

func CreatePet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var decodedRequest = new(createRequest)
		if err := json.NewDecoder(request.Body).Decode(&decodedRequest); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		if decodedRequest.Name == "" {
			http.Error(writer, "name field is required", http.StatusBadRequest)
			return
		}
		if decodedRequest.Specie == "" {
			http.Error(writer, "specie field is required", http.StatusBadRequest)
			return
		}
		if decodedRequest.Sex == "" {
			http.Error(writer, "sex field is required", http.StatusBadRequest)
			return
		}
		/** Para controlar que el campo 'Birthdate' no esté vacío, o contenga su 'zerovalue', es
		necesario implementar la interfaz 'unmarshall' de 'json' dentro de un struct personalizado;
		ergo daremos por supuesto que siempre será facilitado para mantener el ejercicio sencillo,
		ya que el control de errores no ha sido contemplado y/o requerido en el ejercicio propuesto.
		*/
		/*
			if decodedRequest.Birthdate.IsZero() {
				http.Error(writer, "Birthdate field is required", http.StatusBadRequest)
				return
			}
		*/

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
func ListPets() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pageInfo := &models.Pagination{
			OrderBy:   DEFAULT_ORDER_BY,
			PageSize:  DEFAULT_PAGE_SIZE,
			PageToken: 1,
		}

		queryParams := mux.Vars(request)
		if orderBy := queryParams["orderBy"]; orderBy != "" {
			pageInfo.OrderBy = orderBy
		}
		if pageSize := queryParams["pageSize"]; pageSize != "" {
			pageInfo.PageSize = pageSize
		}
		if pageToken := queryParams["pageToken"]; pageToken != "" {
			pageInfo.PageToken, _ = strconv.Atoi(pageToken)
		}
		if totalPages := queryParams["totalPages"]; totalPages != "" {
			pageInfo.TotalPages = totalPages
		}
		if totalItems := queryParams["totalItems"]; totalItems != "" {
			pageInfo.TotalItems = totalItems
		}

		pageInfo, pets, err := databases.ListPets(request.Context(), pageInfo)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		notEncodedResponse := listResponse{
			PageInfo: pageInfo,
			Pets:     pets,
		}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(notEncodedResponse)
	}
}
func ReportPets() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryParams := mux.Vars(request)
		if specie := queryParams["specie"]; specie != "" {
			pageInfo := &models.Pagination{
				OrderBy:        DEFAULT_ORDER_BY,
				PageSize:       "ALL",
				FilterBySpecie: specie,
			}

			_, pets, err := databases.ListPets(request.Context(), pageInfo)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			notEncodedResponse := reportResponse{
				Specie:            "specie not found",
				AverageAge:        "",
				StandardDeviation: "",
			}

			if len(pets) != 0 {
				avgAge, stDev := getAVGAgeAndSTDeviation(pets)

				notEncodedResponse.Specie = pageInfo.FilterBySpecie
				notEncodedResponse.AverageAge = fmt.Sprintf("%f days", avgAge)
				notEncodedResponse.StandardDeviation = fmt.Sprintf("%f days", stDev)
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
