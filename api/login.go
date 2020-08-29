package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
	"github.com/rhizomplatform/log"
)

// @Param email "The user's email."
// @Param password "The user password."
// @Return json jwt Token
func (s *Server) login(c echo.Context) error {
	usr := new(app.User)
	if err := c.Bind(usr); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := usr.Valid(); err != nil {
		return err
	}

	otherUsr := store.FindUserByEmail(s.db, usr.Email.String)
	if otherUsr == nil || otherUsr.Valid() != nil {
		err := fmt.Errorf("This email %s doesn't exists", usr.Email.String)
		log.Error(err)
		return c.String(http.StatusNotFound, err.Error())
	}

	if err := otherUsr.VerifyPwd(*usr); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "This email or password are invalid")
	}

	// set claims
	claims := &JwtClaims{
		UUID:  otherUsr.UUID.Base62(),
		Email: otherUsr.Email.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * jwtExpiresHours).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// generate encoded token
	t, err := token.SignedString([]byte(signedJWTString))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Token{t})
}

// MockJWTToken mocks the jwt token authentication
func (s *Server) MockJWTToken() {

	// set claims
	claims := &JwtClaims{
		UUID:  "6tXaLO1p4I8cLbGkgl8Jgy",
		Email: "adminx201@gmail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * jwtExpiresHours).Unix(),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s.SetCustomContextKey("token", token)
}
