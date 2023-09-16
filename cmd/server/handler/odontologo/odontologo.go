package odontologo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/marinazv/FinalGo/internal/domain/odontologo"
	"github.com/marinazv/FinalGo/pkg/web"
)

type Controlador struct {
	service odontologo.Service
}

func NewControladorOdontologo(service odontologo.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Create a new odontologo
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		odontologo, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologo,
		})

	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Get all odontologos
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologos,
		})
	}
}

// Oontologo godoc
// @Summary odontologo example
// @Description Get odontologo by id
// @Tags odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [get]
func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		odontologo, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologo,
		})
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Update odontologo by id
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

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

		odontologo, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologo,
		})

	}
}

// Oontologo godoc
// @Summary odontologo example
// @Description Delete odontologo by id
// @Tags odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [delete]
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
			"mensaje": "odontologo eliminado",
		})
	}
}

// Patch
func (c *Controlador) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "ID de odontólogo inválido")
			return
		}

		var campos map[string]interface{}
		if err := ctx.BindJSON(&campos); err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Datos de actualización inválidos")
			return
		}

		odontologoActualizado, err := c.service.Patch(ctx, id, campos)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "No se pudo actualizar el odontólogo")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoActualizado,
		})
	}
}
