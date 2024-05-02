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
	occupyScooterHandler         *scooter_handler.OccupyScooterHandler
	releaseScooterHandler        *scooter_handler.ReleaseScooterHandler
	searchScootersHandler        *scooter_handler.SearchScootersHandler
	updateScooterLocationHandler *scooter_handler.UpdateScooterLocationHandler
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

	if err := c.releaseScooterHandler.Handle(scooterId, user.Id); err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}

func (c *ScooterController) UpdateLocation(context *gin.Context) {
	scooterId := uuid.MustParse(context.Param("id"))
	user := context.MustGet("user").(model.User)

	var scooterLocationUpdate dto.ScooterLocationUpdate
	if err := context.BindJSON(&scooterLocationUpdate); err != nil {
		return
	}

	if err := c.updateScooterLocationHandler.Handle(scooterId, user.Id, scooterLocationUpdate); err != nil {
		context.Error(err)
		return
	}

	context.Status(http.StatusNoContent)
}

func NewScooterController(
	occupyScooterHandler *scooter_handler.OccupyScooterHandler,
	releaseScooterHandler *scooter_handler.ReleaseScooterHandler,
	searchScootersHandler *scooter_handler.SearchScootersHandler,
	updateScooterLocationHandler *scooter_handler.UpdateScooterLocationHandler,
) *ScooterController {
	return &ScooterController{
		occupyScooterHandler:         occupyScooterHandler,
		releaseScooterHandler:        releaseScooterHandler,
		searchScootersHandler:        searchScootersHandler,
		updateScooterLocationHandler: updateScooterLocationHandler,
	}
}
