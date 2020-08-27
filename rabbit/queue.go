package rabbit

import "github.com/streadway/amqp"

// Queue stores Queue relational data
type Queue struct {
	Name     string
	Quantity uint16
}

// NewQueue returns a new Queue
func NewQueue(name string) Queue {
	return Queue{
		Name:     name,
		Quantity: 0,
	}
}

// Declare certify that a queue is declared
func (q *Queue) Declare(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		q.Name, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	return err
}

// Publish sends a message to queue
func (q *Queue) Publish(ch *amqp.Channel, msg string) error {
	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

// Consume return a channel to consumes all received messages
func (q *Queue) Consume(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
