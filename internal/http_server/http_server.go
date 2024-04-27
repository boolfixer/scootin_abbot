package http_server

import (
	"github.com/gin-gonic/gin"
	"main/internal/controller"
	"main/internal/middleware"
)

type HttpServer struct {
	router *gin.Engine
}

func (s HttpServer) Serve() {
	err := s.router.Run()

	if err != nil {
		panic(err)
	}
}

func NewHttpServer(scooterController *controller.ScooterController) *HttpServer {
	router := gin.Default()

	public := router.Group("/api/scooters")
	public.GET("/", scooterController.Search)

	protected := router.Group("/api/scooters")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/occupy", scooterController.Occupy)
	protected.POST("/release", scooterController.Release)

	return &HttpServer{router: router}
}
