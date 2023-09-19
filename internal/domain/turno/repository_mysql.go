package turno

import (
	"context"
	"database/sql"
	"log"

	"github.com/marinazv/FinalGo/cmd/server/handler/turno"
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

//CreateByDniAndMatricula creates a new turno 

func (r *repository) CreateByDniAndMatricula(ctx context.Context, request RequestTurnoDniAndMatricula)(any, error){
	statement, err := r.db.Prepare(QueryGetTurnosByDniPaciente)
	if err != err {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		request.Dni,
		request.Matricula,
		request.Descripcion,
		request.FechaHora,
	)

	if err != nil {
		return Turno{}, ErrExec
	}


	return result, nil


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

// Trae turno por dni de paciente que indicamos por query
func (r *repository) GetByDniPaciente(ctx context.Context, dni string) ([]RequestTurnoByDni, error) {
	//queryPaciente := fmt.Sprintf("SELECT id FROM patients WHERE dni = '%s';", dni)

	row, err := r.db.Query(QueryGetTurnosByDniPaciente, dni)
	if err != nil {
		return []RequestTurnoByDni{}, err
	}
	defer row.Close()

	var turnos []RequestTurnoByDni
	for row.Next(){
		var turno RequestTurnoByDni
		err := row.Scan(
		&turno.Descripcion,
		&turno.FechaHora,
		&turno.Dni,
		&turno.NamePaciente,
		&turno.FirstNamePaciente,
		&turno.NameOdontologo,
		&turno.FirstNameOdontologo,
	)

	if err != nil {
		return []RequestTurnoByDni{}, err
	}

	turnos = append(turnos, turno)

}

if err := row.Err(); err != nil {
	return []RequestTurnoByDni{}, err
}

return turnos, nil


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

// Patch
// Patch actualiza parcialmente un turno en la base de datos.
func (r *repository) Patch(ctx context.Context, id int, campos map[string]interface{}) (*Turno, error) {
	// Obtén el turno por su ID utilizando el método GetByID
	turno, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Actualiza los campos del turno con los valores proporcionados en el mapa "campos"
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

	// Actualiza el turno en la base de datos
	_, err = r.db.ExecContext(ctx, QueryUpdateTurno, turno.IdPaciente, turno.IdOdontologo, turno.Descripcion, turno.FechaHora, turno.ID)
	if err != nil {
		return nil, err
	}

	// Retorna el turno actualizado
	return &turno, nil
}
