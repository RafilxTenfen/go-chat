package store

import (
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
)

// FindQueueByName return a queue based on an name
func FindQueueByName(db *gorm.DB, name string) *app.Queue {
	var q app.Queue
	db.Where("name = ?", name).First(&q)

	return &q
}

// InsertQueue inserts a new queue in the database
func InsertQueue(db *gorm.DB, q *app.Queue) error {
	err := db.Create(q).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateQueue update a queue in the database
func UpdateQueue(db *gorm.DB, q *app.Queue) error {
	err := db.Save(q).Error
	if err != nil {
		return err
	}
	return nil
}

// UpsertQueue verifies if exists the queue before inserts, existing it will update the queue pointer ID
func UpsertQueue(db *gorm.DB, q *app.Queue) error {
	if queue := FindQueueByName(db, q.Name.String); queue != nil {
		q.ID = queue.ID
		return UpdateQueue(db, q)
	}
	return InsertQueue(db, q)
}
