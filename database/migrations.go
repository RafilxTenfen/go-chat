package database

import (
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
)

func autoMigrateStructs(db *gorm.DB) {
	db.AutoMigrate(&app.User{})
	db.AutoMigrate(&app.Queue{})
	db.AutoMigrate(&app.Message{})
	db.AutoMigrate(&app.UserQueue{})
	db.Model(&app.UserQueue{}).AddUniqueIndex("idx_unique_user_queue", "user_uuid", "queue_id")
	db.Model(&app.UserQueue{}).AddForeignKey("user_uuid", "users(uuid)", "NO ACTION", "NO ACTION")
	db.Model(&app.UserQueue{}).AddForeignKey("queue_id", "queues(id)", "NO ACTION", "NO ACTION")
	db.Model(&app.Message{}).AddForeignKey("queue_id", "queues(id)", "NO ACTION", "NO ACTION")
}
