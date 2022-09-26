package databases

import (
	"context"
	"database/sql"
	"fmt"
	"rpgmypet/internal/models"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	DEFAULT_PROTOCOOL = "postgres"
)

type PostgresImplementation struct {
	DBConn *sql.DB

	DatabaseConnectionName string
	DatabasePassword       string
	DatabaseProtocool      string
	DatabaseSchema         string
	DatabaseUser           string
}

//* Método para la instanciación de implementaciones PostGreSQL
func NewPostgresImplementation(user, password, connectionName, schema string) (*PostgresImplementation, error) {
	url := buildURL(user, password, connectionName, schema)

	dbConnection, err := sql.Open(DEFAULT_PROTOCOOL, url)
	if err != nil {
		return nil, err
	}

	pgr := &PostgresImplementation{
		DBConn:                 dbConnection,
		DatabaseConnectionName: connectionName,
		DatabasePassword:       password,
		DatabaseProtocool:      DEFAULT_PROTOCOOL,
		DatabaseSchema:         schema,
		DatabaseUser:           user,
	}

	return pgr, nil
}

//* Método privado para construir la URL de conexión a DB.
func buildURL(user, password, connectionName, schema string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, schema, connectionName)
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
	) VALUES ($1,$2,$3,$4,$5)`

	_, err := pgr.DBConn.ExecContext(ctx, querySentence,
		pet.Name,
		pet.Specie,
		pet.Sex,
		pet.Birthdate,
		pet.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

//* Método para obtener todos las entidades "Pet" del sistema CON control de paginación.
func (pgr *PostgresImplementation) ListPets(ctx context.Context, pageInfo *models.Pagination) (*models.Pagination, []*models.Pet, error) {
	querySentence := ""
	if pageInfo.TotalPages == "0" && pageInfo.TotalItems == "0" {
		querySentence = ` SELECT
			count(*) AS total_items
			FROM pets
		`
		if pageInfo.FilterBySpecie != "" {
			querySentence += fmt.Sprintf(" WHERE specie = '%s'", pageInfo.FilterBySpecie)
		}

		rows, err := pgr.DBConn.QueryContext(ctx, querySentence)
		if err != nil {
			pageInfo.TotalPages = "0"
			pageInfo.TotalItems = "0"

			return pageInfo, nil, err
		}

		for rows.Next() {
			if err = rows.Scan(
				&pageInfo.TotalItems,
			); err != nil {
				pageInfo.TotalPages = "0"
				pageInfo.TotalItems = "0"

				return pageInfo, nil, err
			}
		}
		defer rows.Close()
		if err = rows.Err(); err != nil {
			pageInfo.TotalPages = "0"
			pageInfo.TotalItems = "0"

			return pageInfo, nil, err
		}

		totalItems, _ := strconv.Atoi(pageInfo.TotalItems)
		pageSize, _ := strconv.Atoi(pageInfo.PageSize)
		if totalItems%pageSize != 0 {
			pageInfo.TotalPages = fmt.Sprintf("%d", (totalItems/pageSize)+1)
		} else {
			pageInfo.TotalPages = fmt.Sprintf("%d", (totalItems / pageSize))
		}
		pageInfo.PageToken = 1
	}

	querySentence = `SELECT
		name,
		specie,
		sex,
		birthdate,
		id
		FROM pets
	`
	if pageInfo.FilterBySpecie != "" {
		querySentence += fmt.Sprintf(" WHERE specie = '%s'", pageInfo.FilterBySpecie)
	}

	if pageInfo.PageSize == "ALL" {
		querySentence += fmt.Sprintf(" ORDER BY %s LIMIT %s OFFSET 0", pageInfo.OrderBy, pageInfo.PageSize)
	} else {
		pageSize, _ := strconv.Atoi(pageInfo.PageSize)
		querySentence += fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d", pageInfo.OrderBy, pageSize, (pageInfo.PageToken-1)*pageSize)
	}

	rows, err := pgr.DBConn.QueryContext(ctx, querySentence)
	if err != nil {
		pageInfo.TotalPages = "0"
		pageInfo.TotalItems = "0"

		return pageInfo, nil, err
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
			pageInfo.TotalPages = "0"
			pageInfo.TotalItems = "0"

			return pageInfo, nil, err
		}

		pets = append(pets, pet)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		pageInfo.TotalPages = "0"
		pageInfo.TotalItems = "0"

		return pageInfo, nil, err
	}

	pageInfo.PageToken++

	return pageInfo, pets, nil
}
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

	specie := ""
	amount := 0
	for rows.Next() {
		if err = rows.Scan(
			&specie,
			&amount,
		); err != nil {
			return "", err
		}
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return "", err
	}

	return specie, nil
}
