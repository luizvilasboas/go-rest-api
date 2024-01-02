package routes

import (
	"log"
	"net/http"

	"gitlab.com/alura-courses-code/golang/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
