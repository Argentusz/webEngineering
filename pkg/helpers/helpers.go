package helpers

import "strings"

func checkForSpecialSymbols(str string, spaceAllowed bool) bool {
	for _, char := range []rune(str) {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < 'а' || char > 'я') &&
			(char < 'А' || char > 'Я') && (char < '0' || char > '9') &&
			!strings.ContainsRune("!@#^*().,", char) && !(spaceAllowed && char == ' ') {
			return false
		}
	}
	return true
}
