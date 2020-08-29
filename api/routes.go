package api

import (
	"github.com/labstack/echo/v4"
)

// Routes define all the server routes
func (s *Server) Routes() {
	// Public Routes
	s.e.POST("/login", s.login)
	s.e.POST("/add-user", s.addUser)

	// Created restricted auth group
	g := s.e.Group("/api")
	s.RestrictedJWTMiddleware(g)
	s.RestrictedRoutes(g)
}

// RestrictedRoutes define the restricted routes
func (s *Server) RestrictedRoutes(g *echo.Group) {
	s.UserRoutes(g.Group("/user"))
}

// UserRoutes define the Account routes /api/user
func (s *Server) UserRoutes(accountGroup *echo.Group) {
	accountGroup.GET("", s.getAllUser)
	accountGroup.POST("", s.addUser)
	accountGroup.GET("/:userUUID", s.getUser)
	accountGroup.DELETE("/:userUUID", s.deleteUser)
	accountGroup.PUT("/:userUUID", s.updateUser)

	// publish
	accountGroup.POST("/publish/:queueName", s.publishMessage)
	accountGroup.GET("/publish/*", s.publishMessageQuery)

	// get messages
	accountGroup.POST("/messages/:queueName", s.getMessages)
	accountGroup.GET("/messages/*", s.getMessagesQuery)

	accountGroup.Any("/:userUUID/*", echo.NotFoundHandler)
}
