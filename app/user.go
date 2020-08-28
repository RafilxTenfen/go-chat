package app

import (
	"fmt"
	"regexp"

	"github.com/rhizomplatform/log"
	null "github.com/rhizomplatform/pg-null"
)

var (
	regExpEmail    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	errPwdNotMatch = fmt.Errorf("Error the password doesn't match")
)

// User with just email and password
type User struct {
	UUID     null.UUID   `gorm:"primary_key" json:"uuid,omitempty"`
	Email    null.String `gorm:"unique; not null," json:"email,omitempty"`
	Password null.String `gorm:"not null"          json:"password,omitempty"`
}

// NewUser returns a new user structure
func NewUser(email, password string) (*User, error) {
	hash, err := GeneratePwd(email, password)
	if err != nil {
		log.WithError(err).Error("Error on generate Password")
		return nil, err
	}
	return &User{
		UUID:     null.NewID(),
		Email:    null.S(email),
		Password: null.S(hash),
	}, nil
}

// Update updates one user based on another
func (u *User) Update(otherUser User) {
	if otherUser.Email.Valid {
		u.Email = otherUser.Email
	}
	if otherUser.Password.Valid {
		u.Password = otherUser.Password
	}
}

// Valid return nil if the User is valid
func (u User) Valid() error {
	if err := u.ValidEmail(); err != nil {
		return err
	}
	return ValidPwd(u.Password)
}

// ValidEmail return nil if the user email is valid
func (u User) ValidEmail() error {
	if !u.Email.Valid ||
		u.Email.IsNull() ||
		!ValidEmail(u.Email.String) {
		return fmt.Errorf("User Email '%s' is invalid", u.Email.String)
	}
	return nil
}

// GeneratePwd return the Generated password
func (u User) GeneratePwd() (string, error) {
	return GeneratePwd(u.Email.String, u.Password.String)
}

// VerifyPwd return nil if the password has a match
func (u User) VerifyPwd(otherUser User) error {
	match, err := ComparePasswordAndHash(otherUser.Email.String, otherUser.Password.String, u.Password.String)
	if err != nil {
		log.WithError(err).Error("VerifyPwd error on ComparePasswordAndHash")
		return err
	}

	if !match {
		return errPwdNotMatch
	}
	return nil
}

// Clear set the User password to nil
func (u *User) Clear() {
	u.Password = null.String{}
}

// ValidEmail validate the user email
func ValidEmail(email string) bool {
	return regExpEmail.MatchString(email)
}
