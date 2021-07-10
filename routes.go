package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vchicago/overflight-api/api"
)

func SetupRoutes(engine *gin.Engine) {
	overflight := engine.Group("/overflight/v1")
	{
		overflight.GET("/:fac", api.GetLive)
	}
	engine.StaticFile("/overflight", "./static/index.html")
	engine.StaticFile("/overflight/openapi.yaml", "./static/openapi.yaml")

	engine.GET("/overflight/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PONG"})
	})
}
