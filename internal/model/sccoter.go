package model

import "github.com/google/uuid"

type Scooter struct {
	Id        uuid.UUID
	Name      string
	Color     string
	Latitude  int
	Longitude int
}
