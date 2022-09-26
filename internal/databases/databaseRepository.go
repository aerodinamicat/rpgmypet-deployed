package databases

import (
	"context"
	"rpgmypet/internal/models"
)

type DatabaseRepository interface {
	//Common managing
	CloseConnection() error

	//* Métodos estándar CRUD
	//* Create
	InsertPet(ctx context.Context, pet *models.Pet) error

	//* Read
	ListPets(ctx context.Context, pageInfo *models.Pagination) (*models.Pagination, []*models.Pet, error)

	//* Métodos personalizados específicos
	MostCommonPetSpecie(ctx context.Context) (string, error)
}

var dbImplementation DatabaseRepository

func SetDatabaseRepository(dbr DatabaseRepository) {
	dbImplementation = dbr
}

func CloseConnection() error {
	return dbImplementation.CloseConnection()
}
func InsertPet(ctx context.Context, pet *models.Pet) error {
	return dbImplementation.InsertPet(ctx, pet)
}
func ListPets(ctx context.Context, pageInfo *models.Pagination) (*models.Pagination, []*models.Pet, error) {
	return dbImplementation.ListPets(ctx, pageInfo)
}
func MostCommonPetSpecie(ctx context.Context) (string, error) {
	return dbImplementation.MostCommonPetSpecie(ctx)
}
