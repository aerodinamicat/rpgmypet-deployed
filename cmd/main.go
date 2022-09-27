// Package main RPGMyPet API
//
// Así se muestra como hacer una API Rest en Golang.
// Se trata de una vista detallada de las especificaciones del lenguaje.
//
// Terms of Service:
//
// Schemes: http
// Host: localhost:8080
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
	dbConfig := &servers.DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	srv := servers.NewHttpServer(context.Background(), os.Getenv("APP_ENVIRONMENT"), os.Getenv("APP_PORT"), dbConfig)

	srv.Start()
}
