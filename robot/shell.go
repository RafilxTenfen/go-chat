package robot

import (
	"os"

	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/rhizomplatform/log"
	"gopkg.in/abiosoft/ishell.v2"
)

// Shell implements the robot shell commands
func (b *Bot) Shell() *ishell.Shell {
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd{
		Name: "add-queue",
		Help: "add a new queue to the bot queues map.: add-queue {string}",
		Func: b.shellAddQueue,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "exit",
		Help: "exit the program",
		Func: b.shellExit,
	})

	return shell
}

func (b *Bot) shellExit(c *ishell.Context) {
	b.Exit()
	os.Exit(0)
}

func (b *Bot) shellAddQueue(c *ishell.Context) {
	arguments := c.Args
	if len(arguments) != 1 {
		log.Warn("It should add just one queue each time")
	}

	queueName := arguments[0]
	queue := rabbit.NewQueue(queueName)

	if err := queue.Declare(b.channel); err != nil {
		log.WithError(err).Error("error on declare queue")
		return
	}
	if err := b.AddQueueToConsume(queueName); err != nil {
		log.WithError(err).Error("error on add queue")
		return
	}

	log.With(log.F{
		"name": queueName,
	}).Debug("queue added with success")
}
