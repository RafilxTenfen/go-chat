package api

import (
	"fmt"
	"net/http"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/store"
	echo "github.com/labstack/echo/v4"
	"github.com/rhizomplatform/log"
	null "github.com/rhizomplatform/pg-null"
)

// @Param name "The user name."
// @Param email "The user's email."
// @Param password "The user password."
func (s *Server) addUser(c echo.Context) error {
	user := new(app.User)
	if err := c.Bind(user); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := user.Valid(); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := app.NewUser(user.Email.String, user.Password.String)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := store.InsertUser(s.db, user); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	user.Clear()
	return c.JSON(http.StatusOK, user)
}

func (s *Server) getAllUser(c echo.Context) error {
	users := store.GetAllUsers(s.db)

	return c.JSON(http.StatusOK, users)
}

func (s *Server) getUser(c echo.Context) error {
	uuid, err := null.ParseID(c.Param("userUUID"))
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	user := store.FindUserByUUID(s.db, uuid)
	user.Clear()
	if !user.UUID.Valid {
		return c.String(http.StatusBadRequest, fmt.Sprintf("user of uuid '%s' not found", uuid.Base62()))
	}
	user.Clear()
	return c.JSON(http.StatusOK, user)
}

func (s *Server) deleteUser(c echo.Context) error {
	uuid, err := null.ParseID(c.Param("userUUID"))
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := store.DeleteUser(s.db, uuid); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("if exists a user with uuid '%s', it was deleted!", uuid.Base62()))
}

func (s *Server) updateUser(c echo.Context) error {
	uuid, err := null.ParseID(c.Param("userUUID"))
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	user := store.FindUserByUUID(s.db, uuid)
	if err := user.Valid(); err != nil {
		log.With(log.F{
			"user": user,
		}).Error(err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("User of uuid %s doesn't exists", uuid.Base62()))
	}

	userUpdated := new(app.User)
	if err := c.Bind(userUpdated); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.Update(*userUpdated)
	if err := user.Valid(); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := store.UpdateUser(s.db, user); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	user.Clear()
	return c.JSON(http.StatusOK, user)
}

func (s *Server) setUserChatFromRequest(c echo.Context) error {
	jwtClaims, err := getClaimsFromContext(c)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	usr, err := jwtClaims.GetUser(s.db)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := usr.Valid(); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Erro on find user from JWT '%+v'", err))
	}
	s.SetUser(usr)

	return nil
}
