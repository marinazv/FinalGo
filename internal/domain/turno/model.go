package turno


// Turno describes a turno.
type Turno struct {
	ID           int    `json:"id"`
	IdPaciente   int    `json:"id_paciente"`
	IdOdontologo int    `json:"id_odontologo"`
	Descripcion  string `json:"descripcion"`
	FechaHora    string `json:"fecha_hora"`
}

// RequestTurno describes the data needed to create a new turno.
type RequestTurno struct {
	IdPaciente   int    `json:"id_paciente"`
	IdOdontologo int    `json:"id_odontologo"`
	Descripcion  string `json:"descripcion"`
	FechaHora    string `json:"fecha_hora"`
}
