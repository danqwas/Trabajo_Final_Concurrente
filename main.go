package main

import (
	"fmt"
	"lifesafe/controllers"
	"lifesafe/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/consultations/{id}", controllers.GetProductById).Methods("GET")
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
	RegisterRoutes(router)

	// Iniciar el server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
