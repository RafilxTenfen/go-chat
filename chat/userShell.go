package chat

import (
	"os"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/abiosoft/ishell"
	"github.com/jinzhu/gorm"
	"github.com/rhizomplatform/log"
)

// CreateUserShell returns the user shell
func CreateUserShell(usr *app.User, db *gorm.DB) *ishell.Shell {
	shell := ishell.New()

	usrChat, err := NewUserChat(usr, db)
	if err != nil {
		log.WithError(err).Error("error on create user chat")
		return nil
	}

	shell.AddCmd(&ishell.Cmd{
		Name: "publish",
		Help: `Publish a message into a queue "{string(queueName)}" "{string(message)}"`,
		Func: usrChat.shellPublish,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "print",
		Help: `Print all messages from a queue "{string(queueName)}" `,
		Func: usrChat.shellPrintQueue,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "exit",
		Help: "exit the program",
		Func: usrChat.shellExit,
	})

	return shell
}

func (uc *UserChat) shellExit(c *ishell.Context) {
	uc.Exit()
	os.Exit(0)
}

func (uc *UserChat) shellPublish(c *ishell.Context) {
	arguments := c.Args
	if len(arguments) != 2 {
		log.With(log.F{
			"help": c.Cmd.Help,
			"args": arguments,
		}).Warn("Wrong number of parameters")
		return
	}

	queueName := arguments[0]
	msg := arguments[1]
	if err := uc.Publish(queueName, msg); err != nil {
		log.WithError(err).Error("error on publish message")
	}
}

func (uc *UserChat) shellPrintQueue(c *ishell.Context) {
	arguments := c.Args
	if len(arguments) != 1 {
		log.With(log.F{
			"help": c.Cmd.Help,
			"args": arguments,
		}).Warn("Wrong number of parameters")
		return
	}

	queueName := arguments[0]
	msgs, err := uc.Messages(queueName)
	if err != nil {
		log.WithError(err).Error("Error on found messages")
	}

	for _, msg := range msgs {
		log.With(log.F{
			"Msg":  msg.Message.String,
			"Time": msg.CreatedAt,
		}).Info("Print message")
	}
}
