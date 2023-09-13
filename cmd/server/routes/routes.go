package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlerodontologo "github.com/marinazv/FinalGo/cmd/server/handler/odontologo"
	"github.com/marinazv/FinalGo/cmd/server/handler/ping"
	odontologo "github.com/marinazv/FinalGo/internal/domain/Odontologo"
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
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildProductRoutes maps all routes for the odontologos domain.
func (r *router) buildOdontologoRoutes() {
	// Create a new odontologo controller.
	repository := odontologo.NewRepositoryMySql(r.db)
	service := odontologo.NewService(repository)
	controlador := handlerodontologo.NewControladorOdontologo(service)

	r.routerGroup.POST("/odontologos", middleware.Authenticate(),controlador.Create())
	r.routerGroup.GET("/odontologos",middleware.Authenticate(), controlador.GetAll())
	r.routerGroup.GET("/odontologos/:id",middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/odontologos/:id",middleware.Authenticate(), controlador.Update())
	r.routerGroup.DELETE("/odontologos/:id",middleware.Authenticate(), controlador.Delete())

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControladorPing()
	r.routerGroup.GET("/ping", pingController.Ping())

}