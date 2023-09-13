package helpers

import "webEngineering/pkg/models"

func CheckFacultySafety(faculty models.Faculty) bool {
	if faculty.Name == "" || !checkForSpecialSymbols(faculty.Name, true) {
		return false
	}
	if faculty.ExamDate == "" || !checkForSpecialSymbols(faculty.ExamDate, true) {
		return false
	}
	if faculty.ExamAud == "" || !checkForSpecialSymbols(faculty.ExamAud, true) {
		return false
	}
	return true
}
