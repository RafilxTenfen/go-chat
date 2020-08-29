package api

import (
	"net/http"

	"github.com/RafilxTenfen/go-chat/app"
	echo "github.com/labstack/echo/v4"
	"github.com/rhizomplatform/log"
)

func (s *Server) publishMessage(c echo.Context) error {
	queueName := c.Param("queueName")
	msg := new(app.Message)
	if err := c.Bind(msg); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return s.publishMessageIntoQueue(c, queueName, msg.Message.String)
}

func (s *Server) publishMessageIntoQueue(c echo.Context, queueName, msg string) error {
	if msg == "" {
		return c.String(http.StatusBadRequest, "Invalid Message")
	}

	if queueName == "" {
		return c.String(http.StatusBadRequest, "Invalid Queue Name")
	}

	if err := s.setUserChatFromRequest(c); err != nil {
		return err
	}

	return s.usrChat.Publish(queueName, msg)
}

func (s *Server) publishMessageQuery(c echo.Context) error {
	queueName := c.QueryParam("queue")
	msg := c.QueryParam("msg")

	return s.publishMessageIntoQueue(c, queueName, msg)
}
