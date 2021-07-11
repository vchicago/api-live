package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vchicago/api-live/database"
	dbTypes "github.com/vchicago/types/database"
)

func GetFlights(c *gin.Context) {
	var flights []dbTypes.Flights
	database.DB.Where("facility = ?", c.Param("fac")).Find(&flights)

	c.JSON(http.StatusOK, flights)
}

func GetControllers(c *gin.Context) {
	var controllers []dbTypes.OnlineControllers
	database.DB.Where("facility = ?", c.Param("fac")).Find(&controllers)

	c.JSON(http.StatusOK, controllers)
}
