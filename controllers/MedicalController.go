package controllers

import (
	"lifesafe/services"
	"net/http"
)

// Creando una peticion medical
func CreateMedical(w http.ResponseWriter, r *http.Request) {
	services.CreateMedical(w, r)
}

// Obteniendo por ID

func GetConsultationById(w http.ResponseWriter, r *http.Request) {
	services.GetConsultationById(w, r)
}

// Obteniendo todas las consultas
func GetAllConsultation(w http.ResponseWriter, r *http.Request) {
	services.GetAllConsultation(w, r)
}

// Actualizando una consulta
func UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	services.UpdateConsultation(w, r)
}

// Borrando una consulta
func DeletingConsultation(w http.ResponseWriter, r *http.Request) {
	services.DeletingConsultation(w, r)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	services.GetStatus(w, r)
}
