package server

import (
	"fmt"
	"sync"

	"github.com/Arismonx/van-booking-system/config"
	"github.com/Arismonx/van-booking-system/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app  *gin.Engine
	db   database.Database
	conf *config.Config
}

var (
	once           sync.Once
	serverInstance *Server
)

func NewServer(conf *config.Config, db database.Database) *Server {
	app := gin.Default()

	once.Do(func() {
		serverInstance = &Server{
			app:  app,
			db:   db,
			conf: conf,
		}
	})
	return serverInstance
}

func (s *Server) Start() {
	// s.app.Use(gin.Recovery())
	// s.app.Use(gin.Logger())

	// Set up routes
	s.app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Van Booking System API",
		})
	})

	// Start the server
	s.httpListening()
}

func (s *Server) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Run(url); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
