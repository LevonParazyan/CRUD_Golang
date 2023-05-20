package utility

import (
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var passwordPattern = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*]{8,}$`)

func containsDigit(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	symbols := "!@#$%^&*"
	for _, c := range s {
		if strings.ContainsRune(symbols, c) {
			return true
		}
	}
	return false
}

func ValidateEmail(email string) bool {
	return emailPattern.MatchString(email)
}

func ValidatePassword(password string) bool {
	return passwordPattern.MatchString(password) && containsDigit(password) && containsSymbol(password)
}
