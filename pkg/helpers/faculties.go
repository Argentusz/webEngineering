package helpers

import "webEngineering/pkg/models"

func CheckFacultySafety(faculty models.Faculty) bool {
	if !checkForSpecialSymbols(faculty.Name, true) {
		return false
	}
	if !checkForSpecialSymbols(faculty.ExamDate, true) {
		return false
	}
	if !checkForSpecialSymbols(faculty.ExamAud, true) {
		return false
	}
	return true
}
