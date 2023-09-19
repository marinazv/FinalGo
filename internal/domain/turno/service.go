package turno

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetByID(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, campos map[string]interface{}) (*Turno, error)
}

// NewService creates a new turno service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a new turno.
func (s *service) Create(ctx context.Context, requestTurno RequestTurno) (Turno, error) {
	paciente := requestToTurno(requestTurno)
	response, err := s.repository.Create(ctx, paciente)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Turno{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all Turnos.
func (s *service) GetAll(ctx context.Context) ([]Turno, error) {
	turnos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de Turnos", err.Error())
		return []Turno{}, ErrEmptyList
	}

	return turnos, nil
}

// GetByID returns a Turno by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Turno, error) {
	turno, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("log de error en service de turno", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo GetByID")
	}

	return turno, nil
}

// Update updates a Turno.
func (s *service) Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.Update(ctx, turno)
	if err != nil {
		log.Println("log de error en service de turno", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

// Delete deletes a Turno.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado del Turno", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToTurno(requestTurno RequestTurno) Turno {
	var turno Turno
	turno.IdPaciente = requestTurno.IdPaciente
	turno.IdOdontologo = requestTurno.IdOdontologo
	turno.Descripcion = requestTurno.Descripcion
	turno.FechaHora = requestTurno.FechaHora

	return turno
}

// Patch actualiza parcialmente un turno.
func (s *service) Patch(ctx context.Context, id int, campos map[string]interface{}) (*Turno, error) {
	turno, err := s.GetByID(ctx, id)
	if err != nil {
		log.Println("Error al obtener turno en el servicio:", err.Error())
		return nil, err
	}

	// Actualiza los campos del turno con los valores proporcionados en el mapa "campos".
	for campo, valor := range campos {
		switch campo {
		case "id_paciente":
			turno.IdPaciente = valor.(int)
		case "id_odontologo":
			turno.IdOdontologo = valor.(int)
		case "descripcion":
			turno.Descripcion = valor.(string)
		case "fecha_hora":
			fechaAltaStr := valor.(string)
			if err != nil {
				log.Println("Fecha no tiene el formato adecuado")
			}
			turno.FechaHora = fechaAltaStr
		}
	}

	// Actualiza el turno con los nuevos datos.
	turnoActualizado, err := s.repository.Patch(ctx, id, campos) // Pasamos id y campos
	if err != nil {
		log.Println("Error al actualizar turno en el servicio:", err.Error())
		return nil, err
	}

	return turnoActualizado, nil
}
