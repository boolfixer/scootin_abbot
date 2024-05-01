package simulator

import (
	"main/internal/dto"
	"main/internal/http_client"
	"main/internal/model"
	"math/rand"
	"time"
)

type ClientSimulator struct {
	scooterClient http_client.ScooterClient
}

func (s ClientSimulator) Run() {
	// wait to be sure that server bootstrapped and ready to accept requests
	time.Sleep(time.Second * 2)

	for {
		// find first available scooter nearby
		scooter := s.findFirstAvailableScooter()

		// try to occupy scooter or start searching new free scooter
		if !s.scooterClient.Occupy(scooter.Id) {
			continue
		}
		// occupy scooter for some time
		time.Sleep(time.Second * 2)

		// release scooter in random location
		randomLocation := dto.Location{Latitude: randInt(1, 20), Longitude: randInt(1, 20)}
		s.scooterClient.Release(scooter.Id, randomLocation)

		// rest before start new circle
		time.Sleep(time.Second * 2)
	}
}

func (s ClientSimulator) findFirstAvailableScooter() model.Scooter {
	for {
		randomLocation := dto.Location{Latitude: randInt(1, 20), Longitude: randInt(1, 20)}
		scooters := s.scooterClient.Search(randomLocation)

		if len(scooters) > 0 {
			return scooters[0]
		}

		for _, scooter := range scooters {
			// skip occupied scooters
			if scooter.IsOccupied {
				continue
			}

			return scooter
		}
	}
}

func randInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func NewClientSimulator(apiKey string) ClientSimulator {
	return ClientSimulator{
		scooterClient: http_client.NewScooterClient(apiKey),
	}
}
