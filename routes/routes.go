package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/alura-courses-code/golang/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/personalities/all", controllers.HandleAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.HandleGetPersonalityById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
