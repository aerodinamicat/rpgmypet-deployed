package databases

import (
	"context"
	"rpgmypet/internal/models"
)

/** Ésta 'class' representa una 'interface' para la implementación del patrón de diseño 'Repository',
*   que facilita el cambio entre los distintos sistema de base de datos posibles.
 */

type DatabaseRepository interface {
	//Common managing
	CloseConnection() error

	//* Métodos estándar CRUD
	//* Create
	InsertPet(ctx context.Context, pet *models.Pet) error

	//* Read
	ListPets(ctx context.Context, filterBySpecie string) ([]*models.Pet, error)

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
func ListPets(ctx context.Context, filterBySpecie string) ([]*models.Pet, error) {
	return dbImplementation.ListPets(ctx, filterBySpecie)
}
func MostCommonPetSpecie(ctx context.Context) (string, error) {
	return dbImplementation.MostCommonPetSpecie(ctx)
}
