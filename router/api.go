package router

import (
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(h *handler.Handler) *gin.Engine {
	w :=gin.Default()
	/// api 
	h.CreateUser()
	// end 
	return w
}