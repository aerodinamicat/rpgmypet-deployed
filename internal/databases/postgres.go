package databases

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"rpgmypet/internal/models"

	_ "github.com/lib/pq"
)

const (
	DEFAULT_PROTOCOL = "postgres"
	DEV_ENVIRONMENT  = "dev"
	PROD_ENVIRONMENT = "prod"
)

type PostgresImplementation struct {
	DBConn *sql.DB

	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
}

//* Método para crear una nueva instancia de repositorio PostgreSQL.
func NewPostgresImplementation(environment, user, password, host, port, name string) (*PostgresImplementation, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DEFAULT_PROTOCOL, user, password, host, port, name)

	if environment == PROD_ENVIRONMENT {
		url = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, name, host)
	}

	dbConnection, err := sql.Open(DEFAULT_PROTOCOL, url)
	if err != nil {
		return nil, err
	}

	pgr := &PostgresImplementation{
		DBConn:           dbConnection,
		DatabaseUser:     user,
		DatabasePassword: password,
		DatabaseHost:     host,
		DatabaseName:     name,
	}

	return pgr, nil
}

//* Método para el cierre de conexiones a DB
func (pgr *PostgresImplementation) CloseConnection() error {
	return pgr.DBConn.Close()
}

//* Método para añadir nuevas entidades "Pet" al sistema.
func (pgr *PostgresImplementation) InsertPet(ctx context.Context, pet *models.Pet) error {
	querySentence := `INSERT INTO pets (
		name,
		specie,
		sex,
		birthdate,
		id
	) VALUES ($1, $2, $3, $4, $5)`

	_, err := pgr.DBConn.ExecContext(ctx, querySentence,
		pet.Name,
		pet.Specie,
		pet.Sex,
		pet.Birthdate,
		pet.Id,
	)
	if err != nil {
		log.Println("error dbConn")
		return err
	}

	return nil
}

//* Método para obtener todas las entidades "Pet" del sistema. Puede, o no, filtrar por el campo 'specie'.
func (pgr *PostgresImplementation) ListPets(ctx context.Context, filterBySpecie string) ([]*models.Pet, error) {
	querySentence := `SELECT
		name,
		specie,
		sex,
		birthdate,
		id
		FROM pets
	`
	if filterBySpecie != "" {
		querySentence += fmt.Sprintf(" WHERE specie = '%s'", filterBySpecie)
	}

	querySentence += " ORDER BY name asc"

	rows, err := pgr.DBConn.QueryContext(ctx, querySentence)
	if err != nil {
		return nil, err
	}

	var pets []*models.Pet
	for rows.Next() {
		pet := new(models.Pet)
		if err = rows.Scan(
			&pet.Name,
			&pet.Specie,
			&pet.Sex,
			&pet.Birthdate,
			&pet.Id,
		); err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}

//* Método para obtener la especie con mayor índice de sujetos en el sistema.
func (pgr *PostgresImplementation) MostCommonPetSpecie(ctx context.Context) (string, error) {
	querySentence := `SELECT
		specie,
		count(*) as total_pets
		FROM pets
		GROUP BY specie
		ORDER BY total_pets desc
		LIMIT 1
	`

	rows, err := pgr.DBConn.QueryContext(ctx, querySentence)
	if err != nil {
		return "", err
	}

	mostCommonSpecie := ""
	amount := 0
	for rows.Next() {
		if err = rows.Scan(
			&mostCommonSpecie,
			&amount,
		); err != nil {
			return "", err
		}
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return "", err
	}

	return mostCommonSpecie, nil
}
