package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/controller"
	"main/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api/scooters")
	public.GET("/search", controller.ScootersSearch)

	protected := router.Group("/api/scooters")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/occupy", controller.ScooterOccupy)
	protected.POST("/free", controller.ScooterFree)

	return router
}
