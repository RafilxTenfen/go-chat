package store

import (
	"errors"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
	"github.com/rhizomplatform/log"
	null "github.com/rhizomplatform/pg-null"
)

// ExistsUsersQueue return true if the user queue exists
func ExistsUsersQueue(db *gorm.DB, usrQ *app.UserQueue) bool {
	var other app.UserQueue
	result := db.Where("user_uuid = ?", usrQ.UserUUID.String()).Where("queue_id = ?", usrQ.QueueID).Find(&other)

	if other.UserUUID.Valid {
		return true
	}

	err := result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	if err != nil {
		log.WithError(result.Error).Error("error on find user queue")
	}

	return false
}

// DeleteUserQueue deletes all userQueue based on user uuid
func DeleteUserQueue(db *gorm.DB, uuid null.UUID) error {
	return db.Where("user_uuid = ?", uuid).Delete(app.UserQueue{}).Error
}

// InsertUserQueue insert a UserQueue in the database
func InsertUserQueue(db *gorm.DB, usrQ *app.UserQueue) error {
	err := db.Create(usrQ).Error
	if err != nil {
		return err
	}
	return nil
}
