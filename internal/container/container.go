package container

import (
	"go.uber.org/dig"
	"main/internal/controller"
	"main/internal/database"
	"main/internal/repository"
	"main/internal/scooter_handler"
)

var constructors = []interface{}{
	controller.NewScooterController,
	database.NewDBConnection,
	repository.NewScooterRepository,
	repository.NewScooterOccupationRepository,
	scooter_handler.NewSearchScootersHandler,
	scooter_handler.NewOccupyScooterHandler,
	scooter_handler.NewReleaseScooterHandler,
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
