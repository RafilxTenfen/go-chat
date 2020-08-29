package robot

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/rhizomplatform/log"
	"github.com/streadway/amqp"
)

// ConsumeAllQueues consume all the robot queues
func (b *Bot) ConsumeAllQueues() {
	keys := b.queueMap.Keys()
	for i := range keys {
		name := keys[i]
		b.ConsumeQueueByName(name)
	}
}

// ConsumeQueueByName retrieve the queue based on the queue map consumes a queue and start listenning for received msgs
// at this point it assumes that the queue is already been declared
func (b *Bot) ConsumeQueueByName(queueName string) {
	queue, ok := b.queueMap.Load(queueName)
	if !ok {
		log.With(log.F{
			"Queue": queue,
		}).Warn("Error on load queue")
		return
	}

	if err := b.ConsumeQueue(queue); err != nil {
		log.WithError(err).Error("error on consume queue")
	}
}

// ConsumeQueue Consumes a queue and start listenning for received msgs
func (b *Bot) ConsumeQueue(queue *app.Queue) error {
	msgs, err := queue.Consume(b.channel)
	if err != nil {
		return err
	}

	queue.SetConsuming(true)
	if err := store.UpsertQueue(b.db, queue); err != nil {
		return err
	}
	b.queueMap.Store(queue)

	go b.HandleMsgs(queue, msgs)
	return nil
}

// HandleMsgs of an delivery queue channel
func (b *Bot) HandleMsgs(queue *app.Queue, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		b.HandleMsg(queue, d)
	}
}

// HandleMsg handle a single msg
func (b *Bot) HandleMsg(queue *app.Queue, d amqp.Delivery) {
	if IsCommand(d) {
		if err := b.HandleReceivedCommand(d); err != nil {
			log.WithError(err).Error("Error on handle command")
			if err := queue.Publish(b.channel, fmt.Sprintf("Error on handle command: %s", err.Error())); err != nil {
				log.WithError(err).Error("Error on publish message")
			}
		}
	}
	strMessage := string(d.Body)
	log.With(log.F{
		"msg":              strMessage,
		"command":          IsCommand(d),
		"Queue RoutingKey": d.RoutingKey,
	}).Info("Received a message")

	msg := app.NewMessage(strMessage, queue.ID)
	if err := store.InsertMessage(b.db, msg); err != nil {
		log.WithError(err).Error("Error on insert message on database")
		return
	}
}
