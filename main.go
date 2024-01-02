package main

import (
	"fmt"

	"gitlab.com/alura-courses-code/golang/go-rest-api/routes"
)

func main() {
	fmt.Println("Initializing server at http://localhost:8000...")
	routes.HandleRequest()
}
