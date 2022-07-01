package helper

import (
	"lifesafe/database"
	"lifesafe/domain/entities"
)

func CheckIfProductExists(medicalId string) bool {
	var consultation entities.Consultation
	database.Instance.First(&consultation, medicalId)
	if consultation.ID == 0 {
		return false
	}
	return true
}
