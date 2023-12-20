package api

import (
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
				users.POST("/", h.CreateUser)
				users.PUT("/:id", h.UpdateUserById)
				users.GET("/:id", h.GetUserById)
				users.GET("/", h.GetUsers)
				users.DELETE("/:id", h.DeleteUserByID)
				multi := users.Group("/multi")
				{
					multi.POST("/", h.CreateUsers)
					multi.PUT("/", h.UpdateUsers)
					multi.DELETE("/", h.Ping)
				}
			}
		}
	}
	// end
	return w
}
