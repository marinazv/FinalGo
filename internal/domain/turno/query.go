package turno

var (
	QueryInsertTurno = `INSERT INTO my_db.Turnos(id_paciente,id_odontologo,descripcion,fecha_hora)
	VALUES(?,?,?,?)`
	QueryGetAllTurnos = `SELECT id_paciente,id_odontologo,descripcion,fecha_hora 
	FROM my_db.Turnos`
	QueryDeleteTurno  = `DELETE FROM my_db.Turnos WHERE id = ?`
	QueryGetTurnoById = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.Turnos WHERE id = ?`
	QueryUpdateTurno = `UPDATE my_db.Turnos SET id_paciente = ?,id_odontologo = ? ,descripcion = ?,fecha_hora = ?
	WHERE id = ?`
)
