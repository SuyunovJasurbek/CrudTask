package api

import (
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func SetUpApi(h *handler.Handler) *gin.Engine {
	w := gin.Default()
	swagger := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	w.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagger))
	w.GET("/ping", h.Ping)
	
	r := w.Group("/api")
	{
		v := r.Group("/v1")
		{
			users := v.Group("/users")
			{
				// swagger

				users.POST("/", h.CreateUser)
			}
		}
	}
	// end
	return w
}
