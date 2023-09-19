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
)
