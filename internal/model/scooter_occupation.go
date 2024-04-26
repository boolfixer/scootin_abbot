package model

import (
	"github.com/google/uuid"
	"time"
)

type ScooterOccupation struct {
	Id         uuid.UUID
	ScooterId  uuid.UUID
	UserId     uuid.UUID
	OccupiedAt time.Time
	ReleaseAt  time.Time
}
