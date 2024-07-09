package main

import (
	"gin-api-go/database"
	"gin-api-go/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}