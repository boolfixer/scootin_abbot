package container

import (
	"go.uber.org/dig"
	"main/internal/database"
	"main/internal/repository"
)

var constructors = []interface{}{
	database.NewDBConnection,
	repository.NewScooterRepository,
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
