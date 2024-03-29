package turno

import (
	"context"
	"errors"
)

// Errores
var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("turno not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Repository interface {
	Create(ctx context.Context, turno Turno) (Turno, error)
	CreateByDniAndMatricula(ctx context.Context, request RequestTurnoDniAndMatricula)(any, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetByID(ctx context.Context, id int) (Turno, error)
	GetByDniPaciente(ctx context.Context, dni string) ([]RequestTurnoByDni, error)
	Update(ctx context.Context, turno Turno) (Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, campos map[string]interface{}) (*Turno, error)
}
