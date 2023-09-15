package paciente

var (
	QueryInsertPaciente = `INSERT INTO my_db.pacientes(name,first_name,domicilio, dni)
	VALUES(?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, name, first_name, domicilio, dni, fecha_alta
	FROM my_db.pacientes`
	QueryDeletePaciente  = `DELETE FROM my_db.pacientes WHERE id = ?`
	QueryGetPacienteById = `SELECT id, name, first_name, domicilio, dni, fecha_alta
	FROM my_db.pacientes WHERE id = ?`
	QueryUpdatePaciente = `UPDATE my_db.pacientes SET name = ?, first_name = ?, domicilio = ? , dni = ? , fecha_alta = ?
	WHERE id = ?`
)
