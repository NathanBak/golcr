package golcr

import (
	"encoding/json"
	"log"
)

// ClassMember contains information about an individual member in a class.
type ClassMember struct {
	Name           string `json:"name"`
	SpokenName     string `json:"spokenName"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	ID             int64  `json:"id"`
	Phone          string `json:"phone,omitempty"`
	HouseholdPhone string `json:"householdPhone,omitempty"`
	Email          string `json:"email,omitempty"`
	HouseholdEmail string `json:"householdEmail,omitempty"`
	Age            int8   `json:"age,omitempty"`
	Address        string `json:"address,omitempty"`
}

func (m *ClassMember) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
