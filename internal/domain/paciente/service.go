package paciente

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error)
	GetAll(ctx context.Context) ([]Paciente, error)
	GetByID(ctx context.Context, id int) (Paciente, error)
	Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new paciente service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a new paciente.
func (s *service) Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	response, err := s.repository.Create(ctx, paciente)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Paciente{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all pacientes.
func (s *service) GetAll(ctx context.Context) ([]Paciente, error) {
	pacientes, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de pacientes", err.Error())
		return []Paciente{}, ErrEmptyList
	}

	return pacientes, nil
}

// GetByID returns a paciente by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Paciente, error) {
	paciente, err := s.repository.GetByID(ctx, id)
	//paciente.Turnos, err := s.
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo GetByID")
	}

	return paciente, nil
}

// Update updates a paciente.
func (s *service) Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	paciente.ID = id
	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

// Delete deletes a paciente.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de paciente", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToPaciente(requestPaciente RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Name = requestPaciente.Name
	paciente.FirstName = requestPaciente.FirstName
	paciente.Domicilio = requestPaciente.Domicilio
	paciente.Dni = requestPaciente.Dni

	return paciente
}
