package services

import (
	"fmt"
	"net/http"
)

func AddingMedicalConsultance(w http.ResponseWriter, r *http.Request) {
	// this is the code that will be executed when the user clicks the button
	fmt.Fprintf(w, "Hello World!")
}
