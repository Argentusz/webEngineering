package helpers

import (
	"webEngineering/pkg/models"
)

func CheckUniversitySafety(university models.University) bool {
	if len(university.Login) > 32 || !checkForSpecialSymbols(university.Login, false) {
		return false
	}
	if len(university.Password) < 5 || len(university.Password) > 32 {
		return false
	}
	if !checkForSpecialSymbols(university.Name, true) {
		return false
	}
	return true
}
