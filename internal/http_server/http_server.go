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

func NewHttpServer(
	scooterController *controller.ScooterController,
	authMiddleware *middleware.AuthMiddleware,
	errorMiddleware middleware.ErrorMiddleware,
) *HttpServer {
	router := gin.Default()
	router.Use(errorMiddleware.Handle())

	public := router.Group("/api/scooters")
	public.GET("/", scooterController.Search)

	protected := router.Group("/api/scooters")
	protected.Use(authMiddleware.Handle())
	protected.POST("/:id/occupy", scooterController.Occupy)
	protected.POST("/:id/release", scooterController.Release)

	return &HttpServer{router: router}
}
