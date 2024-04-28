package model

import (
	"github.com/google/uuid"
)

type ScooterOccupation struct {
	Id        uuid.UUID
	ScooterId uuid.UUID
	UserId    uuid.UUID
}
