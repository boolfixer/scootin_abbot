package model

import (
	"github.com/google/uuid"
	"time"
)

type ScooterOccupation struct {
	ScooterId  uuid.UUID
	UserId     uuid.UUID
	OccupiedAt time.Time
	ReleaseAt  time.Time
}
