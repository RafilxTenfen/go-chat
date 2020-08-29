package robot

import (
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/database"
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/jinzhu/gorm"
	"github.com/rhizomplatform/log"
	"github.com/streadway/amqp"
)

// Bot represents a Robot to listen and publish at N queues
type Bot struct {
	queueMap   *app.QueueMap
	connection *amqp.Connection
	channel    *amqp.Channel
	db         *gorm.DB
	settings   rabbit.Settings
}

// NewBot returns a new Bot structure
func NewBot(settings rabbit.Settings, queuesName []string) (*Bot, error) {
	conn, ch, err := rabbit.Init(settings.RabbitMqURL)
	if err != nil {
		return nil, err
	}

	db, err := database.DBConnect()
	if err != nil {
		return nil, err
	}

	b := &Bot{
		settings:   settings,
		queueMap:   app.NewQueueMap(),
		connection: conn,
		channel:    ch,
		db:         db,
	}

	for i := range queuesName {
		name := queuesName[i]
		if err := b.AddQueueToConsume(name); err != nil {
			return nil, err
		}
	}

	return b, err
}

// Exit free and close channels
func (b *Bot) Exit() {
	if err := b.channel.Close(); err != nil {
		log.WithError(err).Error("error on close channel")
	}
	if err := b.connection.Close(); err != nil {
		log.WithError(err).Error("error on close channel")
	}
	if err := b.UncossumeQueues(); err != nil {
		log.WithError(err).Error("error on close channel")
	}
}

// AddQueueToConsume will declare, store in queueMap and start to listenning that queue for commands
func (b *Bot) AddQueueToConsume(queueName string) error {
	_, ok := b.queueMap.Load(queueName)
	if ok {
		return app.ErrQueueAlreadyInMap
	}

	queue := app.NewQueue(queueName, false)
	if err := queue.Declare(b.channel); err != nil {
		return err
	}

	return b.ConsumeQueue(queue)
}

// UncossumeQueues set all the queues on the map as consuming to false
func (b *Bot) UncossumeQueues() error {
	keys := b.queueMap.Keys()
	for i := range keys {
		key := keys[i]
		queue, ok := b.queueMap.Load(key)
		if !ok {
			return app.ErrQueueNotFound
		}

		queue.SetConsuming(false)
		if err := store.UpdateQueue(b.db, queue); err != nil {
			return err
		}
	}

	return nil
}
