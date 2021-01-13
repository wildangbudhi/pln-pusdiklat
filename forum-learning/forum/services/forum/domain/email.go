package domain

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

// Email is an Object to handle email format checking
type Email struct {
	value string
}

// NewEmail is an Constructor for Email Object
func NewEmail(email string) (*Email, error) {

	if !validateEmail(email) {
		return nil, fmt.Errorf("Email Format Invalid")
	}

	emailObj := &Email{
		value: email,
	}

	return emailObj, nil
}

// GetValue is a Getter Function for Value
func (obj *Email) GetValue() string {
	return obj.value
}
