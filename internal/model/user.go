package model

import (
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/lllllan02/iam/pkg/utils/stringutil"
	"github.com/lllllan02/iam/pkg/utils/validtor"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model

	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// BeforeCreate Will do the followin:
//
// Verify that the username meets the following requirements:
// not empty, composed only of letters and numbers, unique.
//
// Verify that the format of the email address is correct.
//
// Encrypt the password.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var count int64

	// check username
	if !stringutil.IsAlphaNumberic(u.Username) {
		return errors.WithCode(code.CInvalidUsername, "username %s is invalid", u.Username)
	}
	if tx.Model(&User{}).Where("username = ?", u.Username).Count(&count); count > 0 {
		return errors.WithCode(code.CDuplicateUsername, "username %s already exist", u.Username)
	}

	// check email
	if !validtor.IsEmailValid(u.Email) {
		return errors.WithCode(code.CInvalidEmail, "email %s is invalid", u.Email)
	}
	if tx.Model(&User{}).Where("email = ?", u.Email).Count(&count); count > 0 {
		return errors.WithCode(code.CDuplicaEmail, "email %s already exist", u.Email)
	}

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "encrypt password")
	}
	u.Password = string(hashedPassword)

	return nil
}

// AfterCreate will automatically generate instanceID based on the ID.
func (u *User) AfterCreate(tx *gorm.DB) error {
	if err := tx.Model(&u).Update("instance_id", NewInstanceID("uid-", u.ID)).Error; err != nil {
		return errors.Wrap(err, "generate instance_id")
	}
	return nil
}
