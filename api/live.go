package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vchicago/overflight-api/database"
	dbTypes "github.com/vchicago/types/database"
)

type Flightsv1 struct {
	ID          int       `json:"-"`
	Callsign    string    `json:"callsign"`
	CID         int       `json:"cid"`
	Facility    string    `json:"facility"`
	Latitude    float32   `json:"lat"`
	Longitude   float32   `json:"lon"`
	Groundspeed int       `json:"spd"`
	Heading     int       `json:"hdg"`
	Altitude    int       `json:"alt"`
	Aircraft    string    `json:"type"`
	Departure   string    `json:"dep"`
	Arrival     string    `json:"arr"`
	Route       string    `json:"route"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"lastSeen"`
}

func GetLive(c *gin.Context) {
	var flights []dbTypes.Flights
	database.DB.Where("facility = ?", c.Param("fac")).Find(&flights)
	var retFlights []Flightsv1

	for i := 0; i < len(flights); i++ {
		retFlights = append(retFlights, Flightsv1(flights[i]))
	}

	c.JSON(http.StatusOK, retFlights)
}
