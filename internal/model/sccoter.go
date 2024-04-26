package model

import "github.com/google/uuid"

type Scooter struct {
	Id        uuid.UUID
	Name      string
	Latitude  int
	Longitude int
}
