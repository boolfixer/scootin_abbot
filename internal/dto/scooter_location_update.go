package dto

import "time"

type ScooterLocationUpdate struct {
	Latitude  int       `json:"latitude" binding:"required,numeric,gte=0,lte=20"`
	Longitude int       `json:"longitude" binding:"required,numeric,gte=0,lte=20"`
	Time      time.Time `json:"time" binding:"required"`
}
