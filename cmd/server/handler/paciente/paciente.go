package paciente

import (
	"net/http"
	"strconv"

	// "time"

	"github.com/gin-gonic/gin"
	paciente "github.com/marinazv/FinalGo/internal/domain/paciente"
	"github.com/marinazv/FinalGo/pkg/web"
)

type Controlador struct {
	service paciente.Service
}

func NewControladorPaciente(service paciente.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// paciente godoc
// @Summary paciente example
// @Description Create a new paciente
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.RequestPaciente

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		paciente, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})

	}
}

// paciente godoc
// @Summary paciente example
// @Description Get all pacientes
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pacientes, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": pacientes,
		})
	}
}

// Paciente godoc
// @Summary paciente example
// @Description Get paciente by id
// @Tags paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [get]
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		paciente, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})
	}
}

// paciente godoc
// @Summary paciente example
// @Description Update paciente by id
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.Paciente

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

		paciente, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})

	}
}

// Paciente godoc
// @Summary paciente example
// @Description Delete paciente by id
// @Tags paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [delete]
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
			"mensaje": "paciente eliminado",
		})
	}
}

// Patch
func (c *Controlador) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "ID de Paciente inválido")
			return
		}

		var campos map[string]interface{}
		if err := ctx.BindJSON(&campos); err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Datos de actualización inválidos")
			return
		}

		pacienteActualizado, err := c.service.Patch(ctx, id, campos)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "No se pudo actualizar el paciente")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": pacienteActualizado,
		})
	}
}
