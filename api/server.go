package api

import (
	"net/http"

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

// SetCustomContextKey sets a key in the server's echo.Context
func (s *Server) SetCustomContextKey(key string, val interface{}) {
	s.e.Use(CustomContextMiddleware(key, val))
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

// CustomSkipper implements the Skipper that will
// skip the middleware if a condition is met
func CustomSkipper(e echo.Context) bool {
	v, ok := e.Get("auth").(bool)
	if !ok {
		return false
	}
	return !v
}
