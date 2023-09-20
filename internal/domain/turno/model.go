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

//RequestTurnoByDni describes the data needed to read a turn, pacient and dentist information of a turn
type RequestTurnoByDni struct {
	Descripcion         string `json:"descripcion"`
	FechaHora           string `json:"fecha_hora"`
	Dni                 string `json:"dni"`
	NamePaciente        string `json:"name_paciente"`
	FirstNamePaciente   string `json:"firstname_paciente"`
	NameOdontologo      string `json:"name_odontologo"`
	FirstNameOdontologo string `json:"firstname_odontologo"`
}

type RequestTurnoDniAndMatricula struct {
	Dni         string `json:"dni"`
	Matricula   string `json:"matricula"`
	Descripcion string `json:"descripcion"`
	FechaHora   string `json:"fecha_hora"`
}
