package api

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/rhizomplatform/log"
)

func (s *Server) getMessages(c echo.Context) error {
	queueName := c.Param("queueName")
	return s.getMessagesFromDatabase(c, queueName)
}

func (s *Server) getMessagesQuery(c echo.Context) error {
	queueName := c.QueryParam("queue")
	return s.getMessagesFromDatabase(c, queueName)
}

func (s *Server) getMessagesFromDatabase(c echo.Context, queueName string) error {
	if err := s.setUserChatFromRequest(c); err != nil {
		return err
	}

	msgs, err := s.usrChat.Messages(queueName)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, msgs)
}
