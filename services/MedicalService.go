package services

import (
	"fmt"
	"net/http"
)

func AddingMedicalConsultance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding Medical Consultance")
	fmt.Fprintf(w, "Adding Medical Consultance")
}
