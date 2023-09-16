package turno

var (
	QueryInsertTurno = `INSERT INTO my_db.Turnos(id_paciente,id_odontologo,descripcion,fecha_hora)
	VALUES(?,?,?,?)`
	QueryGetAllTurnos = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora 
	FROM my_db.Turnos`
	QueryDeleteTurno  = `DELETE FROM my_db.Turnos WHERE id = ?`
	QueryGetTurnoById = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.Turnos WHERE id = ?`
	QueryGetTurnosByIdPaciente = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.Turnos WHERE id_paciente = ?`
	QueryUpdateTurno = `UPDATE my_db.Turnos SET id_paciente = ?,id_odontologo = ? ,descripcion = ?,fecha_hora = ?
	WHERE id = ?`
	QueryGetTurnosByDniPaciente = `SELECT id, id_paciente, descripcion, fecha_hora FROM my_db.Turnos INNER JOIN my_db.Pacientes ON my_db.Turnos.id_paciente = my_db.Pacientes.id WHERE my_db.Pacientes.dni = ?`
)
