package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// RootMiddleware define all the middleware used
func (s *Server) RootMiddleware() {
	s.e.Logger = LogrusAdapter()
	s.e.Logger.SetLevel(log.DEBUG)
	s.e.Use(Logger())
	s.e.Use(middleware.Recover())
}

// RestrictedJWTMiddleware set JWT middleware for N echo groups
func (s Server) RestrictedJWTMiddleware(groups ...*echo.Group) {
	for i := range groups {
		g := groups[i]
		g.Use(middleware.JWTWithConfig(JwtMiddlewareConfig()))
	}
}

// CustomContextMiddleware sets pairs 'key:value' in the default echo.Context
func CustomContextMiddleware(key string, val interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(key, val)
			return next(c)
		}
	}
}
