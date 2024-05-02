package simulator

import (
	"github.com/google/uuid"
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
		// 1. find first available scooter nearby
		scooter := s.findFirstAvailableScooter()

		// 2. try to occupy scooter or start searching new free scooter
		if !s.scooterClient.Occupy(scooter.Id) {
			continue
		}

		time.Sleep(time.Second * 2)
		// 3. send first location update
		if !s.sendLocationUpdate(scooter.Id) {
			break
		}

		time.Sleep(time.Second * 2)
		// 4. send second location update
		if !s.sendLocationUpdate(scooter.Id) {
			break
		}

		// 5. release scooter
		s.scooterClient.Release(scooter.Id)

		// 6. rest before start new circle
		time.Sleep(time.Second * 2)
	}
}

func (s ClientSimulator) findFirstAvailableScooter() model.Scooter {
	for {
		latitude, longitude := generateRandomLocationCoordinates()
		randomLocation := dto.Location{Latitude: latitude, Longitude: longitude}
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

func (s ClientSimulator) sendLocationUpdate(scooterId uuid.UUID) bool {
	latitude, longitude := generateRandomLocationCoordinates()
	scooterLocationUpdate := dto.ScooterLocationUpdate{Latitude: latitude, Longitude: longitude, Time: time.Now()}

	return s.scooterClient.UpdateScooterLocation(scooterId, scooterLocationUpdate)
}

func generateRandomLocationCoordinates() (latitude int, longitude int) {
	return randInt(1, 20), randInt(1, 20)
}

func randInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func NewClientSimulator(apiKey string) ClientSimulator {
	return ClientSimulator{
		scooterClient: http_client.NewScooterClient(apiKey),
	}
}
