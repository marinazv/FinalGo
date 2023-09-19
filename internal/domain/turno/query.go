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
	QueryGetTurnosByDniPaciente = `SELECT t.fecha_hora, t.descripcion, p.dni AS paciente, p.name AS paciente, p.first_name AS paciente, 
o.name AS odontologo, o.first_name AS odontologo
FROM turnos t
JOIN pacientes p ON t.id_paciente = p.id
JOIN odontologos o ON t.id_odontologo = o.id
WHERE p.dni = ?;`
)
