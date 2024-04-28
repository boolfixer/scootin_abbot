package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/model"
	"main/internal/scooter_handler"
	"net/http"
	"strconv"
)

type ScooterController struct {
	occupyScooterHandler  *scooter_handler.OccupyScooterHandler
	releaseScooterHandler *scooter_handler.ReleaseScooterHandler
	searchScootersHandler *scooter_handler.SearchScootersHandler
}

func (c *ScooterController) Search(context *gin.Context) {
	latitude, err := strconv.Atoi(context.Query("latitude"))

	if err != nil {
		panic(err)
	}

	longitude, err := strconv.Atoi(context.Query("longitude"))

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, c.searchScootersHandler.Handle(latitude, longitude))
}

func (c *ScooterController) Occupy(context *gin.Context) {
	scooterId := uuid.MustParse(context.Param("id"))
	user := context.MustGet("user").(model.User)

	c.occupyScooterHandler.Handle(scooterId, user.Id)

	context.Status(http.StatusCreated)
}

func (c *ScooterController) Release(context *gin.Context) {
	scooterId := uuid.MustParse(context.Param("id"))
	user := context.MustGet("user").(model.User)

	var scooterLocation dto.Location

	if err := context.BindJSON(&scooterLocation); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.releaseScooterHandler.Handle(scooterId, user.Id, scooterLocation)
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
