package helpers

import "webEngineering/pkg/models"

func CheckStudentSafety(student models.Student) bool {
	return checkForSpecialSymbols(student.Name, true)
}
