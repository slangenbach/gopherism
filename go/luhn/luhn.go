package luhn

import (
	"strings"
	"unicode"
)

// Check if id string is valid per definition of Luhn algorithm
func Valid(id string) bool {
	var count int
	trimmedId := strings.ReplaceAll(id, " ", "")
	// first digit is doubled
	isDoubled := len(trimmedId)%2 == 0

	if len(trimmedId) <= 1 {
		return false
	}

	for _, char := range trimmedId {

		if !unicode.IsDigit(char) {
			return false
		}

		// convert char to digit
		digit := int(char - '0')

		if isDoubled {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}

		}

		count += digit
		isDoubled = !isDoubled

	}

	return count%10 == 0
}
