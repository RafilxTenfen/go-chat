package app

import (
	"github.com/jinzhu/gorm"
	null "github.com/rhizomplatform/pg-null"
)

// Message defines a queue message
type Message struct {
	gorm.Model
	Message null.String `gorm:"not null" json:"message,omitempty"`
	QueueID uint        `gorm:"not null" json:"-"`
}

// NewMessage returns a new Message
func NewMessage(msg string, qID uint) *Message {
	return &Message{
		Message: null.S(msg),
		QueueID: qID,
	}
}
