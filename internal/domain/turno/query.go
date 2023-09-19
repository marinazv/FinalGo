package turno

var (
	QueryInsertTurno = `INSERT INTO my_db.turnos(id_paciente,id_odontologo,descripcion,fecha_hora)
	VALUES(?,?,?,?)`
	QueryGetDniPacientes = `SELECT dni, FROM my_db.pacientes WHERE id = ?`
	QueryGetMatriculaOdontologo = `SELECT matricula FROM my_db.odontologos WHERE id = ?`
	QueryGetAllTurnos = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos`
	QueryDeleteTurno  = `DELETE FROM my_db.turnos WHERE id = ?`
	QueryGetTurnoById = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos WHERE id = ?`
	QueryGetTurnosByIdPaciente = `SELECT id, id_paciente,id_odontologo,descripcion,fecha_hora
	FROM my_db.turnos WHERE id_paciente = ?`
	QueryUpdateTurno = `UPDATE my_db.turnos SET id_paciente = ?,id_odontologo = ? ,descripcion = ?,fecha_hora = ?
	WHERE id = ?`
	QueryGetTurnosByDniPaciente = `INSERT INTO turnos (id_paciente, id_odontologo, descripcion, fecha_hora)
	SELECT (SELECT id FROM pacientes WHERE dni = ?) AS id_paciente,
	(SELECT id FROM odontologos WHERE matricula = ?) AS id_odontologo, ? AS descripcion, ? AS fecha_hora;`
)
