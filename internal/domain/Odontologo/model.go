package odontologo

import (
	"github.com/marinazv/FinalGo/internal/domain/turno"
)

// Odontologo describes a odontologo.
type Odontologo struct {
	ID        int           `json:"id"`
	FirstName      string        `json:"name"`
	LastName string        `json:"first_name"`
	Matricula string        `json:"matricula"`
	Turnos    []turno.Turno `json:"turnos"`
}

// RequestOdontologo describes the data needed to create a new odontologo.
type RequestOdontologo struct {
	FirstName      string `json:"name"`
	LastName string `json:"first_name"`
	Matricula string `json:"matricula"`
}
