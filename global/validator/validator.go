package validator

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(mailAddress string) bool {
	_, err := mail.ParseAddress(mailAddress)
	return err == nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	passwordPattern := `^[A-Za-z\d]{8,}$`
	_, err := regexp.MatchString(passwordPattern, password)
	return err == nil
}

func ValidatePhone(phoneNumber string) bool {
	if len(phoneNumber) > 13 || len(phoneNumber) < 12 || phoneNumber[:1] != "0" {
		return false
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' || char == '+' {
			return false
		}
	}
	return true
}
