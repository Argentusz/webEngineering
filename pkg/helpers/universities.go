package helpers

import (
	"strings"
	"webEngineering/pkg/models"
)

func CheckUniversitySafety(university models.University) bool {
	if len(university.Login) > 32 || !checkForSpecialSymbols(university.Login, false) {
		return false
	}
	if len(university.Password) < 5 || len(university.Password) > 32 || !checkForSpecialSymbols(university.Password, false) {
		return false
	}
	if !checkForSpecialSymbols(university.Name, true) {
		return false
	}
	return true
}

func checkForSpecialSymbols(str string, spaceAllowed bool) bool {
	for _, char := range []rune(str) {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') &&
			!strings.ContainsRune("!@#^*().,", char) && !(spaceAllowed && char == ' ') {
			return false
		}
	}
	return true
}
