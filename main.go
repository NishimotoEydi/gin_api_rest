package main

import (
	"projetosAlura/gin_api/database"
	"projetosAlura/gin_api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
