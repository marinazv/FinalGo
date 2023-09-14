package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
}

func NewControladorPing() *Controlador {
	return &Controlador{}
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (c *Controlador) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
