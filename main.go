package main

import (
	"fmt"
	"lifesafe/controllers"
	"lifesafe/database"
	"lifesafe/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Variables Globales

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/consultations/{id}", controllers.GetConsultationById).Methods("GET")
	router.HandleFunc("/api/consultations", controllers.CreateMedical).Methods("POST")
	router.HandleFunc("/api/consultations", controllers.GetAllConsultation).Methods("GET")
	router.HandleFunc("/api/consultations/{id}", controllers.UpdateConsultation).Methods("PUT")
	router.HandleFunc("/api/consultations/{id}", controllers.DeletingConsultation).Methods("DELETE")
}

func main() {
	// Cargando configuraciones del config.json
	LoadAppConfig()

	// Inicializar Base de datos
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Inicializar las rutas (Controllers)
	router := mux.NewRouter().StrictSlash(true)

	// Registrar las rutas

	utils.Initialize()
	RegisterRoutes(router)
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
	fmt.Println("Welcome to MedicalRecordApp! 😇")
	// Iniciar el server
}
