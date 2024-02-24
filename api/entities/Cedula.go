package entities

import "github.com/pborman/uuid"

type Cedula struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewCedula() *Cedula {
	cedula := Cedula{
		ID: uuid.New(),
	}
	return &cedula
}
