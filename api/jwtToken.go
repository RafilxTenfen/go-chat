package api

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	null "github.com/rhizomplatform/pg-null"
)

const (
	jwtExpiresHours = 168
	signedJWTString = "api-char-server-key"
)

var (
	// ErrInvalidToken error unsed when a token is invalid
	ErrInvalidToken = fmt.Errorf("JWT token is invalid, please add a valid token to header: Bearer 'token'")
)

// Token is a user credential for authentication
type Token struct {
	JWTToken string `json:"token"`
}

// JwtClaims claims for store user data
type JwtClaims struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// JwtMiddlewareConfig returns a new middleware jwt token
func JwtMiddlewareConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte(signedJWTString),
		ContextKey: "token",
	}
}

func getClaimsFromContext(c echo.Context) (*JwtClaims, error) {
	token, ok := c.Get("token").(*jwt.Token)
	if !ok {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// GetUser returns the user based on jwtClaims structure
func (j *JwtClaims) GetUser(db *gorm.DB) (*app.User, error) {
	uuid, err := null.ParseID(j.UUID)
	if err != nil {
		return nil, err
	}

	return store.FindUserByUUID(db, uuid), nil
}
