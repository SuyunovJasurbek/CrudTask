package api

import (
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(h *handler.Handler) *gin.Engine {
	w := gin.New()
	config := cors.DefaultConfig()
	config.AllowHeaders = append(
		config.AllowHeaders,
		`Content-Type,
		Content-Length,
		Accept-Encoding,
		X-CSRF-Token,
		Authorization,
		accept,
		origin,
		Cache-Control,
		X-Requested-With`,
	)
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	w.Use(cors.New(config))
	w.Use(MaxAllowed(100))
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
					multi.DELETE("/", h.DeleteUsers)
				}
			}
		}
	}
	// end
	return w
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}