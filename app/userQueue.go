package app

import (
	null "github.com/rhizomplatform/pg-null"
	"github.com/streadway/amqp"
)

// UserQueue store queue names for User
type UserQueue struct {
	UserUUID null.UUID `gorm:"not null"    json:"-"`
	QueueID  uint      `gorm:"not null"    json:"-"`
	Queue    *Queue    `                   json:"queue,omitempty"`
}

// NewUserQueue returns a new User Queue
func NewUserQueue(usrUUID null.UUID, q *Queue) *UserQueue {
	return &UserQueue{
		UserUUID: usrUUID,
		QueueID:  q.ID,
		Queue:    q,
	}
}

// Declare certify that a queue is declared
func (usrQ *UserQueue) Declare(ch *amqp.Channel) error {
	return usrQ.Queue.Declare(ch)
}

// Publish sends a message to queue
func (usrQ *UserQueue) Publish(ch *amqp.Channel, msg string) error {
	return usrQ.Queue.Publish(ch, msg)
}

// Consume return a channel to consumes all received messages
func (usrQ *UserQueue) Consume(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	return usrQ.Queue.Consume(ch)
}
