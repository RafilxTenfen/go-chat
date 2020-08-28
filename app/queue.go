package app

import (
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/jinzhu/gorm"
	null "github.com/rhizomplatform/pg-null"
	"github.com/streadway/amqp"
)

// Queue store queue names
type Queue struct {
	gorm.Model
	Name      null.String `gorm:"unique; not null"  json:"name,omitempty"`
	Consuming null.Bool   `                         json:"bot_consuming,omitempty"`
	Messages  []Message   `                         json:"messages,omitempty"`
}

// NewQueue returns a new queue
func NewQueue(name string, consuming bool) *Queue {
	return &Queue{
		Name:      null.S(name),
		Consuming: null.B(consuming),
	}
}

// SetConsuming updates the consuming property
func (q *Queue) SetConsuming(b bool) {
	q.Consuming.Set(b)
}

// Declare certify that a queue is declared
func (q *Queue) Declare(ch *amqp.Channel) error {
	return rabbit.Declare(ch, q.Name.String)
}

// Publish sends a message to queue
func (q *Queue) Publish(ch *amqp.Channel, msg string) error {
	return rabbit.Publish(ch, q.Name.String, msg)
}

// Consume return a channel to consumes all received messages
func (q *Queue) Consume(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	return rabbit.Consume(ch, q.Name.String)
}
