package odontologo

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new odontologo service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a new odontologo.
func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Odontologo{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all odontologos.
func (s *service) GetAll(ctx context.Context) ([]Odontologo, error) {
	odontologos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return []Odontologo{}, ErrEmptyList
	}

	return odontologos, nil
}

// GetByID returns a odontologo by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Odontologo, error) {
	odontologo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("log de error en service de odontolo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo GetByID")
	}

	return odontologo, nil
}

// Update updates a odontologo.
func (s *service) Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	odontologo.ID = id
	response, err := s.repository.Update(ctx, odontologo)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

// Delete deletes a odontologo.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de odontologo", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.FirstName = requestOdontologo.FirstName
	odontologo.LastName = requestOdontologo.LastName
	odontologo.Matricula = requestOdontologo.Matricula

	return odontologo
}
