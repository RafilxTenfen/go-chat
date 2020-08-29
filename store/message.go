package store

import (
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
)

// FindMessagesFromUserQueue return all messages from that queue
func FindMessagesFromUserQueue(db *gorm.DB, usrQ *app.UserQueue, limit uint16) (msgs []app.Message) {
	db.
		Table("messages").
		Joins("JOIN queues ON messages.queue_id = queues.id").
		Joins("JOIN user_queues ON user_queues.queue_id = queues.id").
		Where("user_queues.user_uuid = ?", usrQ.UserUUID.String()).
		Where("user_queues.queue_id = ?", usrQ.QueueID).
		Order("created_at desc").
		Limit(limit).
		Find(&msgs)

	return msgs
}

// InsertMessageData inserts a new message in the database using strings
func InsertMessageData(db *gorm.DB, queueName, msg string) error {
	q := FindQueueByName(db, queueName)
	if q == nil {
		q = app.NewQueue(queueName, false)
		if err := InsertQueue(db, q); err != nil {
			return err
		}
	}

	message := app.NewMessage(msg, q.ID)
	return InsertMessage(db, message)
}

// InsertMessage inserts a new message in the database
func InsertMessage(db *gorm.DB, msg *app.Message) error {
	err := db.Create(msg).Error
	if err != nil {
		return err
	}

	return nil
}
