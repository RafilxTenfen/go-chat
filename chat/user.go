package chat

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/jinzhu/gorm"
	"github.com/rhizomplatform/log"
	"github.com/streadway/amqp"
)

// UserChat wraps the user with chat function
type UserChat struct {
	user       *app.User
	connection *amqp.Connection
	channel    *amqp.Channel
	db         *gorm.DB
	settings   rabbit.Settings
}

// NewUserChatStructure return a new user chat structure
func NewUserChatStructure(usr *app.User, db *gorm.DB, conn *amqp.Connection, ch *amqp.Channel, st rabbit.Settings) *UserChat {
	return &UserChat{
		user:       usr,
		connection: conn,
		channel:    ch,
		db:         db,
		settings:   st,
	}
}

// NewUserChat returns a UserChat structure based on an user
func NewUserChat(usr *app.User, db *gorm.DB) (*UserChat, error) {
	st := rabbit.LoadSettingsFromEnv()

	conn, ch, err := rabbit.Init(st.RabbitMqURL)
	if err != nil {
		log.WithError(err).Error("Error on init rabbitMQ")
		return nil, err
	}

	return NewUserChatStructure(usr, db, conn, ch, st), nil
}

// SetUser sets the user into userChat structure
func (uc *UserChat) SetUser(user *app.User) {
	uc.user = user
}

// Exit closes the user channel
func (uc *UserChat) Exit() {
	if err := uc.channel.Close(); err != nil {
		log.WithError(err).Error("error on close user channel")
	}
	if err := uc.connection.Close(); err != nil {
		log.WithError(err).Error("error on close user connection")
	}
	if err := uc.db.Close(); err != nil {
		log.WithError(err).Error("error on close the database connection")
	}
}

// Publish a message in certain queue
func (uc *UserChat) Publish(queueName, msg string) error {
	log.With(log.F{
		"queueName": queueName,
		"msg":       msg,
	}).Debug("Publishing msg")

	queue := store.FindQueueByName(uc.db, queueName)
	if queue.ID == 0 {
		queue = app.NewQueue(queueName, false)
		if err := store.InsertQueue(uc.db, queue); err != nil {
			return err
		}
		if err := queue.Declare(uc.channel); err != nil {
			return err
		}
	}
	usrQ := app.NewUserQueue(uc.user.UUID, queue)
	if !store.ExistsUsersQueue(uc.db, usrQ) {
		if err := store.InsertUserQueue(uc.db, usrQ); err != nil {
			log.WithError(err).Error("error on insert user queue")
			return err
		}
	}
	return usrQ.Publish(uc.channel, msg)
}

// Messages return all the messages
func (uc *UserChat) Messages(queueName string) ([]app.Message, error) {
	queue := store.FindQueueByName(uc.db, queueName)
	if queue.ID == 0 {
		return []app.Message{}, fmt.Errorf("Queue '%s' not found", queueName)
	}

	usrQ := app.NewUserQueue(uc.user.UUID, queue)
	if !store.ExistsUsersQueue(uc.db, usrQ) {
		return []app.Message{}, fmt.Errorf("The user '%s' is not registered on queue '%s' found", uc.user.UUID.String(), queueName)
	}

	return store.FindMessagesFromUserQueue(uc.db, usrQ, uc.settings.QuantityMessageQueue), nil
}
