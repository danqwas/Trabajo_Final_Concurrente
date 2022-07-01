package services

import (
	"encoding/json"
	"lifesafe/database"
	"lifesafe/domain/entities"
	"lifesafe/helper"
	"net/http"

	"github.com/gorilla/mux"
)

// Crea una consulta
func CreateMedical(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var consultation entities.Consultation
	json.NewDecoder(r.Body).Decode(&consultation)
	database.Instance.Create(&consultation)
	json.NewEncoder(w).Encode(consultation)
}

// Obtiene una consulta
func GetConsultationById(w http.ResponseWriter, r *http.Request) {
	consultationId := mux.Vars(r)["id"]
	if helper.CheckIfProductExists(consultationId) == false {
		json.NewEncoder(w).Encode("Consultation Not Found!")
		return
	}
	var consultation entities.Consultation
	database.Instance.First(&consultation, consultationId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultation)
}

// Obtiene todas las consultas
func GetAllConsultation(w http.ResponseWriter, r *http.Request) {
	var consultation []entities.Consultation
	database.Instance.Find(&consultation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consultation)
}

func UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	consultationId := mux.Vars(r)["id"]
	if helper.CheckIfProductExists(consultationId) == false {
		json.NewEncoder(w).Encode("Consultation Not Found!")
		return
	}
	var consultation entities.Consultation
	database.Instance.First(&consultation, consultationId)
	json.NewDecoder(r.Body).Decode(&consultation)
	database.Instance.Save(&consultation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultation)
}

func DeletingConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	consultationId := mux.Vars(r)["id"]
	if helper.CheckIfProductExists(consultationId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Consultation Not Found!")
		return
	}
	var consultation entities.Consultation
	database.Instance.Delete(&consultation, consultationId)
	json.NewEncoder(w).Encode("consultation Deleted Successfully!")
}
