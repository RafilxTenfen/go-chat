package api

import (
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/chat"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Server basic server structure
type Server struct {
	e       *echo.Echo
	db      *gorm.DB
	usrChat *chat.UserChat
}

// NewServer returns a new server pointer
func NewServer(e *echo.Echo, db *gorm.DB, userChat *chat.UserChat) *Server {
	return &Server{
		e:       e,
		db:      db,
		usrChat: userChat,
	}
}

// SetUser sets the user into the UserChat structure
func (s *Server) SetUser(user *app.User) {
	s.usrChat.SetUser(user)
}
