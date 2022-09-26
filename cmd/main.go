// Package main RPGMyPet API
//
// Así se muestra como hacer una API Rest en Golang.
// Se trata de una vista detallada de las especificaciones del lenguaje.
//
// Terms of Service:
//
// Schemes: http
// Host: localhost:7050
// BasePath: /
// Version: 1.0.0
// License: MIT http://opensource.org/licenses/MIT
// Contact: Saam Aerodinamicat <aerodinamicat@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"os"
	"rpgmypet/internal/servers"
)

func main() {
	config := &servers.Config{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}

	dbConfig := &servers.DBConfig{
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
		ConnectionName: os.Getenv("CLOUD_SQL_CONNECTION_NAME"),
		Schema:         os.Getenv("DB_SCHEMA"),
	}

	srv := servers.NewHttpServer(context.Background(), config, dbConfig)

	srv.Start()
}
