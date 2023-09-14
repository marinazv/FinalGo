package paciente

// Paciente describes a paciente.
type Paciente struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Domicilio string `json:"domicilio"`
	Dni       string `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}

// RequestPaciente describes the data needed to create a new paciente.
type RequestPaciente struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Domicilio string `json:"domicilio"`
	Dni       string `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}
