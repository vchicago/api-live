package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vchicago/api-live/middleware"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(appenv string) *Server {
	server := Server{}

	engine := gin.New()
	engine.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	engine.Use(cors.New(corsConfig))
	engine.Use(middleware.Logger)
	server.engine = engine
	engine.LoadHTMLGlob("static/*")

	SetupRoutes(engine)

	return &server
}
