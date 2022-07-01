package controllers

import (
	"lifesafe/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMedicalController(router *mux.Router) {
	router.HandleFunc("/", services.AddingMedicalConsultance)
}

func MedicalController() {
	router := mux.NewRouter().StrictSlash(true)
	CreateMedicalController(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
