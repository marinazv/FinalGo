package paciente

var (
	QueryInsertPaciente = `INSERT INTO my_db.pacientes(name,first_name,domicilio, dni)
	VALUES(?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, name, first_name, domicilio, dni, fecha_alta
	FROM my_db.Pacientes`
	QueryDeletePaciente  = `DELETE FROM my_db.Pacientes WHERE id = ?`
	QueryGetPacienteById = `SELECT id, name, first_name, domicilio, dni, fecha_alta
	FROM my_db.Pacientes WHERE id = ?`
	QueryUpdatePaciente = `UPDATE my_db.Pacientes SET name = ?, first_name = ?, domicilio = ? , dni = ? , fecha_alta = ?
	WHERE id = ?`
)
