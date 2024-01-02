package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/alura-courses-code/golang/go-rest-api/database"
	"gitlab.com/alura-courses-code/golang/go-rest-api/models"
)

func HandleAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func HandleGetPersonalityById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
