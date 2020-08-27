package robot

import (
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/streadway/amqp"
)

// Bot represents a Robot to listen and publish at N queues
type Bot struct {
	queueMap   *rabbit.QueueMap
	connection *amqp.Connection
	chanel     *amqp.Channel
	settings   Settings
}

// NewBot returns a new Bot structure
func NewBot(settings Settings, queuesName []string) (*Bot, error) {

	conn, ch, err := rabbit.Init(settings.RabbitMqURL)
	if err != nil {
		return nil, err
	}

	queues := rabbit.NewQueueMap()
	for i := range queuesName {
		name := queuesName[i]
		queue := rabbit.NewQueue(name)

		if err := queue.Declare(ch); err != nil {
			return nil, err
		}
		queues.Store(queue)
	}

	return &Bot{
		settings:   settings,
		queueMap:   queues,
		connection: conn,
		chanel:     ch,
	}, err
}
