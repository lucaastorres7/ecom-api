package main

import (
	"fmt"
	"log"

	"github.com/lucaastorres7/ecom-api/cmd/api"
	"github.com/lucaastorres7/ecom-api/config"
	"github.com/lucaastorres7/ecom-api/database"
)

func main() {
	db, err := database.NewPostgresStorage(database.DBConfig{
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Host:     config.Envs.DBHost,
		Port:     config.Envs.DBPort,
		Database: config.Envs.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	database.InitStorage(db)

	port := fmt.Sprintf(":%s", config.Envs.Port)
	server := api.NewAPIServer(port, db)
	if err := server.Run(); err != nil {
		log.Fatal()
	}
}
