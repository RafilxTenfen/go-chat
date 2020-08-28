package store

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
	null "github.com/rhizomplatform/pg-null"
)

// FindUser return a user based on an email or uuid
func FindUser(db *gorm.DB, user app.User) *app.User {
	if user.ValidEmail() == nil {
		return FindUserByEmail(db, user.Email.String)
	}

	return FindUserByUUID(db, user.UUID)
}

// FindUserByEmail return a user based on an email
func FindUserByEmail(db *gorm.DB, email string) *app.User {
	var usr app.User
	db.Where("email = ?", email).First(&usr)

	return &usr
}

// FindUserByUUID return a user based on an uuid
func FindUserByUUID(db *gorm.DB, uuid null.UUID) *app.User {
	var usr app.User
	db.Where("uuid = ?", uuid).First(&usr)

	return &usr
}

// InsertUser inserts a new user in the database
func InsertUser(db *gorm.DB, usr *app.User) error {
	user := FindUserByEmail(db, usr.Email.String)
	if err := user.Valid(); err == nil {
		return fmt.Errorf("user with the email %s already exists", usr.Email.String)
	}

	err := db.Create(usr).Error
	if err != nil {
		return err
	}
	usr.Clear()

	return nil
}

// GetUser return an user based on ID
func GetUser(db *gorm.DB, id uint) (user app.User) {
	db.Find(&user, id)
	user.Clear()
	return
}

// GetAllUsers return all users from the database
func GetAllUsers(db *gorm.DB) (users []app.User) {
	db.Find(&users)
	clearUsers(users)
	return
}

func clearUsers(users []app.User) {
	for i := range users {
		users[i].Clear()
	}
}

// DeleteUser deletes a user based on uuid
func DeleteUser(db *gorm.DB, uuid null.UUID) error {
	return db.Where("uuid = ?", uuid).Delete(app.User{}).Error
}

// UpdateUser update a user
func UpdateUser(db *gorm.DB, usr *app.User) error {
	return db.Save(usr).Error
}
