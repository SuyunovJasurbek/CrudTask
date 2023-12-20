package handler

import (
	"net/http"

	"github.com/SuyunovJasurbek/CrudTask/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
// Ping
// @Summary  Ping
// @Description  Ping
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router        /ping	[get]
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}
