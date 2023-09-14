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
	GetAll(ctx context.Context) ([]Turno, error)
	GetByID(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, turno Turno) (Turno, error)
	Delete(ctx context.Context, id int) error
}
