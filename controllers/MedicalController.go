package controllers

import (
	"encoding/json"
	"lifesafe/database"
	"lifesafe/domain/entities"
	"net/http"

	"github.com/gorilla/mux"
)

// Creando una peticion medical
func CreateMedical(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var consultation entities.Consultation
	json.NewDecoder(r.Body).Decode(&consultation)
	database.Instance.Create(&consultation)
	json.NewEncoder(w).Encode(consultation)
}

// Obteniendo por ID
func checkIfProductExists(medicalId string) bool {
	var consultation entities.Consultation
	database.Instance.First(&consultation, medicalId)
	if consultation.ID == 0 {
		return false
	}
	return true
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	consultationId := mux.Vars(r)["id"]
	if checkIfProductExists(consultationId) == false {
		json.NewEncoder(w).Encode("Consultation Not Found!")
		return
	}
	var consultation entities.Consultation
	database.Instance.First(&consultation, consultationId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultation)
}

// Obteniendo todas las consultas
func GetAllConsultation(w http.ResponseWriter, r *http.Request) {
	var consultation []entities.Consultation
	database.Instance.Find(&consultation)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consultation)
}

// Actualizando una consulta
func UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	consultationId := mux.Vars(r)["id"]
	if checkIfProductExists(consultationId) == false {
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

// Borrando una consulta
func DeletingConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	consultationId := mux.Vars(r)["id"]
	if checkIfProductExists(consultationId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Consultation Not Found!")
		return
	}
	var consultation entities.Consultation
	database.Instance.Delete(&consultation, consultationId)
	json.NewEncoder(w).Encode("consultation Deleted Successfully!")
}
