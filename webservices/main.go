package main

import (
	"log"
	"net/http"
	"webservices/database"
	"webservices/product"

	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
