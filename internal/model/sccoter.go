package model

import (
	"github.com/google/uuid"
	"time"
)

type Scooter struct {
	Id                uuid.UUID
	Name              string
	Latitude          int
	Longitude         int
	LocationUpdatedAt time.Time
	IsOccupied        bool
}
