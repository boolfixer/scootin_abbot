package routes

import (
	"github.com/gin-gonic/gin"
	"main/internal/controllers"
	"main/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api/scooters")
	public.GET("/search", controllers.ScootersSearch)

	protected := router.Group("/api/scooters")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/occupy", controllers.ScooterOccupy)
	protected.POST("/free", controllers.ScooterFree)

	return router
}
