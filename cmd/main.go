package main

import (
	"log"

	"github.com/steevedevops/first-go-api/cmd/api"
	"github.com/steevedevops/first-go-api/db"
)

func main() {

	dbConn := db.NewDBserver("webmaster", "pgsql.dev", "localhost", "firstgoapi", 5433)
	db, err := dbConn.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Nao foi possivel conectar com o servidor")
	}
}
