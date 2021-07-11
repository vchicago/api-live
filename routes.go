package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vchicago/api-live/api"
)

func SetupRoutes(engine *gin.Engine) {
	live := engine.Group("/v1/live")
	{
		live.GET("/flights/:fac", api.GetFlights)
		live.GET("/controllers/:fac", api.GetControllers)
	}
	engine.StaticFile("/v1/live", "./static/index.html")
	engine.StaticFile("/v1/live/openapi.yaml", "./static/openapi.yaml")

	engine.GET("/v1/live/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PONG"})
	})
}
