package odontologo

var (
	QueryInsertOdontologo = `INSERT INTO my_db.odontologos(name,first_name,matricula)
	VALUES(?,?,?)`
	QueryGetAllOdontologos = `SELECT id, name, first_name, matricula 
	FROM my_db.odontologos`
	QueryDeleteOdontologo  = `DELETE FROM my_db.odontologos WHERE id = ?`
	QueryDeleteTurno = `DELETE FROM my_db.turnos WHERE id_odontologo = ?`
	QueryGetOdontologoById = `SELECT id, name, first_name, matricula
	FROM my_db.odontologos WHERE id = ?`
	QueryUpdateOdontologo = `UPDATE my_db.odontologos SET name = ?, first_name = ?, matricula = ?
	WHERE id = ?`
)
