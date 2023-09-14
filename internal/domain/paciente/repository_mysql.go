package paciente

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

// Create creates a new paciente.
func (r *repository) Create(ctx context.Context, paciente Paciente) (Paciente, error) {

	statement, err := r.db.Prepare(QueryInsertPaciente)

	if err != nil {
		return Paciente{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Name,
		paciente.FirstName,
		paciente.Domicilio,
		paciente.Dni,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Paciente{}, ErrLastId
	}

	paciente.ID = int(lastId)

	return paciente, nil
}

// GetAll returns all pacientes.
func (r *repository) GetAll(ctx context.Context) ([]Paciente, error) {
	rows, err := r.db.Query(QueryGetAllPacientes)
	if err != nil {
		return []Paciente{}, err
	}

	defer rows.Close()

	var pacientes []Paciente

	for rows.Next() {
		var paciente Paciente
		err := rows.Scan(
			&paciente.ID,
			&paciente.Name,
			&paciente.FirstName,
			&paciente.Domicilio,
			&paciente.Dni,
		)
		if err != nil {
			return []Paciente{}, err
		}

		pacientes = append(pacientes, paciente)
	}

	if err := rows.Err(); err != nil {
		return []Paciente{}, err
	}

	return pacientes, nil
}

// GetByID returns a paciente by its ID.
func (r *repository) GetByID(ctx context.Context, id int) (Paciente, error) {
	row := r.db.QueryRow(QueryGetPacienteById, id)

	var paciente Paciente
	err := row.Scan(
		&paciente.ID,
		&paciente.Name,
		&paciente.FirstName,
		&paciente.Domicilio,
		&paciente.Dni,
		&paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, err
	}

	return paciente, nil
}

// Update updates a paciente.
func (r *repository) Update(ctx context.Context, paciente Paciente) (Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdatePaciente)
	if err != nil {
		return Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Name,
		paciente.FirstName,
		paciente.Domicilio,
		paciente.Dni,
		paciente.FechaAlta,
		paciente.ID,
	)

	if err != nil {
		return Paciente{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Paciente{}, err
	}

	if rowsAffected < 1 {
		return Paciente{}, ErrNotFound
	}

	return paciente, nil
}

// Delete deletes a paciente.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePaciente, id)
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
