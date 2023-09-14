package turno

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// NewRepositoryMySql creates a new repository.
func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new turno.
func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare(QueryInsertTurno)

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.IdPaciente,
		turno.IdOdontologo,
		turno.Descripcion,
		turno.FechaHora,
	)

	if err != nil {
		return Turno{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Turno{}, ErrLastId
	}

	turno.ID = int(lastId)

	return turno, nil
}

// GetAll returns all turnos.
func (r *repository) GetAll(ctx context.Context) ([]Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return []Turno{}, err
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno
		err := rows.Scan(
			&turno.ID,
			&turno.IdPaciente,
			&turno.IdOdontologo,
			&turno.Descripcion,
			&turno.FechaHora,
		)
		if err != nil {
			return []Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []Turno{}, err
	}

	return turnos, nil
}

// GetByID returns a turno by its ID.
func (r *repository) GetByID(ctx context.Context, id int) (Turno, error) {
	row := r.db.QueryRow(QueryGetTurnoById, id)

	var turno Turno
	err := row.Scan(
		&turno.ID,
		&turno.IdPaciente,
		&turno.IdOdontologo,
		&turno.Descripcion,
		&turno.FechaHora,
	)

	if err != nil {
		return Turno{}, err
	}

	return turno, nil
}

// Update updates a turno.
func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.IdPaciente,
		turno.IdOdontologo,
		turno.Descripcion,
		turno.FechaHora,
		turno.ID,
	)

	if err != nil {
		return Turno{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, ErrNotFound
	}

	return turno, nil
}

// Delete deletes a turno.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurno, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil

}
