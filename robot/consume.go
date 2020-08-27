package robot

import (
	"github.com/RafilxTenfen/go-chat/rabbit"
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
func (b *Bot) ConsumeQueue(queue rabbit.Queue) error {
	msgs, err := queue.Consume(b.channel)
	if err != nil {
		return err
	}

	go b.HandleMsgs(queue.Name, msgs)
	return nil
}

// HandleMsgs of an delivery queue channel
func (b *Bot) HandleMsgs(queueName string, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		b.HandleMsg(queueName, d)
	}
}

// HandleMsg handle a single msg
func (b *Bot) HandleMsg(queueName string, d amqp.Delivery) {
	if IsCommand(d) {
		b.HandleReceivedCommand(d)
	}
	b.queueMap.Add(queueName)
	log.With(log.F{
		"msg":               string(d.Body),
		"command":           IsCommand(d),
		"Regexp Find":       string(regCommand.Find(d.Body)),
		"Regexp Find Value": string(regCommandValue.Find(d.Body)),
		"Queue RoutingKey":  d.RoutingKey,
		"Queue Type":        d.Type,
	}).Debug("Received a message")
}
