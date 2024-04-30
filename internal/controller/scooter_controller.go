package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/model"
	"main/internal/scooter_handler"
	"net/http"
)

type ScooterController struct {
	occupyScooterHandler  *scooter_handler.OccupyScooterHandler
	releaseScooterHandler *scooter_handler.ReleaseScooterHandler
	searchScootersHandler *scooter_handler.SearchScootersHandler
}

func (c *ScooterController) Search(context *gin.Context) {
	var userLocation dto.Location
	if err := context.BindQuery(&userLocation); err != nil {
		return
	}

	context.JSON(http.StatusOK, c.searchScootersHandler.Handle(userLocation))
}

func (c *ScooterController) Occupy(context *gin.Context) {
	scooterId := uuid.MustParse(context.Param("id"))
	user := context.MustGet("user").(model.User)

	if err := c.occupyScooterHandler.Handle(scooterId, user.Id); err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusCreated)
}

func (c *ScooterController) Release(context *gin.Context) {
	scooterId := uuid.MustParse(context.Param("id"))
	user := context.MustGet("user").(model.User)

	var scooterLocation dto.Location
	if err := context.BindJSON(&scooterLocation); err != nil {
		return
	}

	if err := c.releaseScooterHandler.Handle(scooterId, user.Id, scooterLocation); err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}

func NewScooterController(
	occupyScooterHandler *scooter_handler.OccupyScooterHandler,
	releaseScooterHandler *scooter_handler.ReleaseScooterHandler,
	searchScootersHandler *scooter_handler.SearchScootersHandler,
) *ScooterController {
	return &ScooterController{
		occupyScooterHandler:  occupyScooterHandler,
		releaseScooterHandler: releaseScooterHandler,
		searchScootersHandler: searchScootersHandler,
	}
}
