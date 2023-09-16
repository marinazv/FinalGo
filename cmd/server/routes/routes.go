package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlerodontologo "github.com/marinazv/FinalGo/cmd/server/handler/odontologo"
	handlerPaciente "github.com/marinazv/FinalGo/cmd/server/handler/paciente"
	"github.com/marinazv/FinalGo/cmd/server/handler/ping"
	handlerTurno "github.com/marinazv/FinalGo/cmd/server/handler/turno"
	odontologo "github.com/marinazv/FinalGo/internal/domain/odontologo"
	paciente "github.com/marinazv/FinalGo/internal/domain/paciente"
	"github.com/marinazv/FinalGo/internal/domain/turno"
	"github.com/marinazv/FinalGo/pkg/middleware"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildOdontologoRoutes()
	r.buildPacienteRoutes()
	r.buildTurnoRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildOdontologoRoutes maps all routes for the odontologos domain.
func (r *router) buildOdontologoRoutes() {
	// Create a new odontologo controller.
	repository := odontologo.NewRepositoryMySql(r.db)
	service := odontologo.NewService(repository)
	controlador := handlerodontologo.NewControladorOdontologo(service)

	r.routerGroup.POST("/odontologos", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/odontologos", middleware.Authenticate(), controlador.GetAll())
	r.routerGroup.GET("/odontologos/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/odontologos/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.DELETE("/odontologos/:id", middleware.Authenticate(), controlador.Delete())
	r.routerGroup.PATCH("/odontologos/:id", middleware.Authenticate(), controlador.Patch())

}

// buildPacienteRoutes maps all routes for the pacientes domain.
func (r *router) buildPacienteRoutes() {
	// Create a new paciente controller.
	repository := paciente.NewRepositoryMySql(r.db)
	service := paciente.NewService(repository)
	controlador := handlerPaciente.NewControladorPaciente(service)

	r.routerGroup.POST("/pacientes", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/pacientes", middleware.Authenticate(), controlador.GetAll())
	r.routerGroup.GET("/pacientes/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/pacientes/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.DELETE("/pacientes/:id", middleware.Authenticate(), controlador.Delete())
	r.routerGroup.PATCH("/pacientes/:id", middleware.Authenticate(), controlador.Patch())

}

// buildTurnoRoutes maps all routes for the turno domain.
func (r *router) buildTurnoRoutes() {
	// Create a new turno controller.
	repository := turno.NewRepositoryMySql(r.db)
	service := turno.NewService(repository)
	controlador := handlerTurno.NewControladorTurno(service)

	r.routerGroup.POST("/turnos", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/turnos", middleware.Authenticate(), controlador.GetAll())
	r.routerGroup.GET("/turnos/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/turnos/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.DELETE("/turnos/:id", middleware.Authenticate(), controlador.Delete())
	r.routerGroup.PATCH("/turnos/:id", middleware.Authenticate(), controlador.Patch())

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControladorPing()
	r.routerGroup.GET("/ping", pingController.Ping())

}
