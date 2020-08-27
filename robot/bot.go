package robot

import (
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/rhizomplatform/log"
	"github.com/streadway/amqp"
)

// Bot represents a Robot to listen and publish at N queues
type Bot struct {
	queueMap   *rabbit.QueueMap
	connection *amqp.Connection
	channel    *amqp.Channel
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
		channel:    ch,
	}, err
}

// Exit free and close channels
func (b *Bot) Exit() {
	if err := b.channel.Close(); err != nil {
		log.WithError(err).Error("error on close channel")
	}
	if err := b.connection.Close(); err != nil {
		log.WithError(err).Error("error on close channel")
	}
}

// AddQueueToConsume will declare, store in queueMap and start to listenning that queue for commands
func (b *Bot) AddQueueToConsume(queueName string) error {
	_, ok := b.queueMap.Load(queueName)
	if ok {
		return rabbit.ErrQueueAlreadyInMap
	}

	queue := rabbit.NewQueue(queueName)
	if err := queue.Declare(b.channel); err != nil {
		return err
	}

	b.queueMap.Store(queue)
	return b.ConsumeQueue(queue)
}
