package turno

var (
	QueryInsertTurno = `INSERT INTO my_db.turnos(id_paciente,id_odontologo,descripcion,fecha_hora)
	VALUES(?,?,?,?)`
	QueryGetAllTurnos = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos`
	QueryDeleteTurno  = `DELETE FROM my_db.turnos WHERE id = ?`
	QueryGetTurnoById = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos WHERE id = ?`
	QueryGetTurnosByIdPaciente = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos WHERE id_paciente = ?`
	QueryUpdateTurno = `UPDATE my_db.turnos SET id_paciente = ?,id_odontologo = ? ,descripcion = ?,fecha_hora = ?
	WHERE id = ?`
	QueryGetTurnosByDniPaciente = `SELECT descripcion, fecha_hora, my_db.Pacientes.dni, my_db.Pacientes.name, my_db.Pacientes.first_name, my_db.Odontologos.name, my_db.Odontologos.first_name FROM my_db.Turnos JOIN my_db.Odontologos ON my_db.Turnos.id_odontologo = my_db.Odontologos.id JOIN my_db.Pacientes ON my_db.Turnos.id_paciente = my_db.Pacientes.id WHERE my_db.Pacientes.dni = ?`
)
