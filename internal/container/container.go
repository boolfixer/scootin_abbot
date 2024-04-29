package container

import (
	"go.uber.org/dig"
	"main/internal/controller"
	"main/internal/database"
	"main/internal/http_server"
	"main/internal/middleware"
	"main/internal/repository"
	"main/internal/scooter_handler"
)

var constructors = []interface{}{
	controller.NewScooterController,
	database.NewDBConnection,
	middleware.NewAuthMiddleware,
	middleware.NewErrorMiddleware,
	repository.NewScooterRepository,
	repository.NewScooterOccupationRepository,
	repository.NewUserRepository,
	scooter_handler.NewSearchScootersHandler,
	scooter_handler.NewOccupyScooterHandler,
	scooter_handler.NewReleaseScooterHandler,
	http_server.NewHttpServer,
}

func Bootstrap() *dig.Container {
	container := dig.New()

	for _, constructor := range constructors {
		err := container.Provide(constructor)

		if err != nil {
			panic(err)
		}
	}

	return container
}
