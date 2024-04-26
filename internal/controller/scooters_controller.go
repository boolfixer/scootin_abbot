package controller

import (
	"github.com/gin-gonic/gin"
	"main/internal/scooter_handler"
)

type ScooterController struct {
	occupyScooterHandler  scooter_handler.OccupyScooterHandler
	releaseScooterHandler scooter_handler.ReleaseScooterHandler
	searchScootersHandler scooter_handler.SearchScootersHandler
}

func (c ScooterController) Search(context *gin.Context) {

}

func (c ScooterController) Occupy(context *gin.Context) {

}

func (c ScooterController) Release(context *gin.Context) {

}

func NewScooterController(
	occupyScooterHandler scooter_handler.OccupyScooterHandler,
	releaseScooterHandler scooter_handler.ReleaseScooterHandler,
	searchScootersHandler scooter_handler.SearchScootersHandler,
) ScooterController {
	return ScooterController{
		occupyScooterHandler:  occupyScooterHandler,
		releaseScooterHandler: releaseScooterHandler,
		searchScootersHandler: searchScootersHandler,
	}
}
