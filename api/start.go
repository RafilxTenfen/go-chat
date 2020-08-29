package api

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/chat"
	"github.com/RafilxTenfen/go-chat/database"
	"github.com/labstack/echo/v4"
	"github.com/rhizomplatform/log"
)

// Run the server
func Run(port uint16) error {
	e := echo.New()
	db, err := database.DBConnect()
	if err != nil {
		log.Error(err)
		return err
	}

	usrChat, err := chat.NewUserChat(nil, db)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Error(err)
		}
		usrChat.Exit()
	}()

	server := NewServer(e, db, usrChat)
	server.Routes()
	server.RootMiddleware()

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	return nil
}
