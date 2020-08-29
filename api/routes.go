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
	s.PublishRoutes(g.Group("/publish"))
	s.MessageRoutes(g.Group("/messages"))
}

// UserRoutes define the User routes /api/user
func (s *Server) UserRoutes(accountGroup *echo.Group) {
	accountGroup.GET("", s.getAllUser)
	accountGroup.POST("", s.addUser)
	accountGroup.GET("/:userUUID", s.getUser)
	accountGroup.DELETE("/:userUUID", s.deleteUser)
	accountGroup.PUT("/:userUUID", s.updateUser)

	accountGroup.Any("/:userUUID/*", echo.NotFoundHandler)
}

// PublishRoutes define the Publish routes /api/publish
func (s *Server) PublishRoutes(accountGroup *echo.Group) {
	accountGroup.POST("/:queueName", s.publishMessage)
	accountGroup.GET("/*", s.publishMessageQuery)
}

// MessageRoutes define the Message routes /api/messages
func (s *Server) MessageRoutes(accountGroup *echo.Group) {
	accountGroup.POST("/:queueName", s.getMessages)
	accountGroup.GET("/*", s.getMessagesQuery)
}
