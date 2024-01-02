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

func HandleCreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func HandleDeletePersonality(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personality
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func HandleUpdatePersonality(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
