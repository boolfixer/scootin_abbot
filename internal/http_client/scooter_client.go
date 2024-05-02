package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"main/internal/dto"
	"main/internal/model"
	"net/http"
	"strconv"
	"time"
)

type ScooterClient struct {
	apiKey     string
	httpClient http.Client
}

func (c ScooterClient) Search(location dto.Location) []model.Scooter {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/scooters", nil)

	q := req.URL.Query()
	q.Add("latitude", strconv.Itoa(location.Latitude))
	q.Add("longitude", strconv.Itoa(location.Longitude))

	req.URL.RawQuery = q.Encode()
	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic("Failed to send request.")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var scooters []model.Scooter
	if err := json.Unmarshal(responseBody, &scooters); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return scooters
}

func (c ScooterClient) Occupy(scooterId uuid.UUID) (scooterOccupied bool) {
	url := fmt.Sprintf("http://localhost:8080/api/scooters/%s/occupy", scooterId)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic("Failed to send request.")
	}

	return resp.StatusCode == http.StatusCreated
}

func (c ScooterClient) UpdateScooterLocation(scooterId uuid.UUID, scooterLocationUpdate dto.ScooterLocationUpdate) bool {
	marshalled, _ := json.Marshal(scooterLocationUpdate)

	url := fmt.Sprintf("http://localhost:8080/api/scooters/%s/update-location", scooterId)
	req, _ := http.NewRequest(http.MethodPatch, url, bytes.NewReader(marshalled))

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic("Failed to send request.")
	}

	return resp.StatusCode == http.StatusNoContent
}

func (c ScooterClient) Release(scooterId uuid.UUID) (scooterReleased bool) {
	url := fmt.Sprintf("http://localhost:8080/api/scooters/%s/release", scooterId)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic("Failed to send request.")
	}

	return resp.StatusCode != http.StatusNoContent
}

func NewScooterClient(apiKey string) ScooterClient {
	return ScooterClient{
		apiKey:     apiKey,
		httpClient: http.Client{Timeout: time.Second * 5},
	}
}
