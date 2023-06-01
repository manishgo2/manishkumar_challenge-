package main

import (
	"fmt"
	"regexp"
	"strings"
)

// CardValidation checks the validity of a credit card number.
// It returns true if the card number is valid, and false otherwise.
func CardValidation(CardNumber string) bool {
	// Regular expression pattern to validate credit card number format
	pattern := `^(4|5|6)\d{3}(-?\d{4}){3}$`

	// Check if the card number matches the pattern
	match, _ := regexp.MatchString(pattern, CardNumber)
	if !match {
		return false
	}

	// Remove any hyphens from the input
	CardNumber = strings.ReplaceAll(CardNumber, "-", "")

	// Check for the consecutive repeation of the digits
	for i := 0; i < len(CardNumber)-3; i++ {
		if CardNumber[i] == CardNumber[i+1] && CardNumber[i] == CardNumber[i+2] && CardNumber[i] == CardNumber[i+3] {
			return false
		}
	}

	return true
}

func main() {
	// Read the number of credit card numbers to validate
	var n int
	fmt.Scanln(&n)

	// Read each credit card number and perform validation
	CardNumbers := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&CardNumbers[i])
	}

	// Validate each credit card number
	for _, CardNumber := range CardNumbers {
		if CardValidation(CardNumber) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
