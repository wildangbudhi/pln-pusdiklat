package domain

import (
	"fmt"
	"regexp"
)

func validateUUID(uuid string) bool {
	re := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return re.MatchString(uuid)
}

// UUID is an Object to handle uuid format checking
type UUID struct {
	value string
}

// NewUUID is an Constructor for UUID Object
func NewUUID(uuid string) (*UUID, error) {

	if !validateUUID(uuid) {
		return nil, fmt.Errorf("UUID Format Invalid")
	}

	uuidObj := &UUID{
		value: uuid,
	}

	return uuidObj, nil

}

// GetValue is a Getter Function for Value
func (obj *UUID) GetValue() string {
	return obj.value
}
