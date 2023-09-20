package turno

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	turno "github.com/marinazv/FinalGo/internal/domain/turno"
	"github.com/marinazv/FinalGo/pkg/web"
)

type Controlador struct {
	service turno.Service
}

func NewControladorTurno(service turno.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// turno godoc
// @Summary turno example
// @Description Create a new turno
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turno, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusCreated, gin.H{
			"data": turno,
		})

	}
}

// turno godoc
// @Summary turno example
// @Description Create a new turno by Dni and Matricula
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnosByDni [post]
func (c *Controlador)CreateTurno() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		var request turno.RequestTurnoDniAndMatricula

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turno, err := c.service.CreateByDniAndMatricula(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusCreated, gin.H{
			"message": turno,
		})

	}
}

// turno godoc
// @Summary turno example
// @Description Get all turnos
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /turnos [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turnos,
		})
	}
}

// Turno godoc
// @Summary turno example
// @Description Get turno by id
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [get]
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		turno, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})
	}
}

// Turno godoc
// @Summary turno example
// @Description Get turno by pacienteDni
// @Tags turno
// @Param dni path string "dni del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnosPaciente/dni [get]
func (c *Controlador) GetByDniPaciente() gin.HandlerFunc{
	return func(ctx *gin.Context){
		dni := ctx.Query("dni")
		turno, err := c.service.GetByDniPaciente(ctx, dni)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})
	}
}

// turno godoc
// @Summary turno example
// @Description Update turno by id
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		turno, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})

	}
}

// Turno godoc
// @Summary turno example
// @Description Delete turno by id
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [delete]
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "turno eliminado",
		})
	}
}



//Patch
func (c *Controlador) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "ID de Turno inválido")
			return
		}

		var campos map[string]interface{}
		if err := ctx.BindJSON(&campos); err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Datos de actualización inválidos")
			return
		}

		turnoActualizado, err := c.service.Patch(ctx, id, campos)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "No se pudo actualizar el turno")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turnoActualizado,
		})
	}
}