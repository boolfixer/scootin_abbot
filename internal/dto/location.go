package dto

type Location struct {
	Latitude  int `json:"latitude" form:"latitude" binding:"required,numeric,gte=0,lte=20"`
	Longitude int `json:"longitude" form:"longitude" binding:"required,numeric,gte=0,lte=20"`
}
