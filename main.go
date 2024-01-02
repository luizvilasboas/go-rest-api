package main

import (
	"fmt"

	"gitlab.com/alura-courses-code/golang/go-rest-api/database"
	"gitlab.com/alura-courses-code/golang/go-rest-api/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Initializing server at http://localhost:8000...")
	routes.HandleRequest()
}
