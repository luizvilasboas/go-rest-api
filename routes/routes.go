package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/alura-courses-code/golang/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/personality/all", controllers.HandleAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personality/{id}", controllers.HandleGetPersonalityById).Methods("GET")
	r.HandleFunc("/api/personality", controllers.HandleCreatePersonality).Methods("POST")
	r.HandleFunc("/api/personality/{id}", controllers.HandleDeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personality/{id}", controllers.HandleUpdatePersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
