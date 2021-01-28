package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// UUID is an Object to handle uuid format checking
type UUID struct {
	value string
}

// NewUUID is an Constructor for UUID Object generating UUID from Google UUID
func NewUUID() *UUID {
	return &UUID{
		value: uuid.New().String(),
	}
}

// NewUUIDFromString is an Constructor for UUID Object from String
func NewUUIDFromString(uuid string) (*UUID, error) {

	uuidObj := &UUID{
		value: uuid,
	}

	if !uuidObj.validateUUID() {
		return nil, fmt.Errorf("UUID Format Invalid")
	}

	return uuidObj, nil

}

// GetValue is a Getter Function for Value
func (obj *UUID) GetValue() string {
	return obj.value
}

func (obj *UUID) validateUUID() bool {
	_, err := uuid.Parse(obj.value)

	if err != nil {
		return false
	}

	return true

}
