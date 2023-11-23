package model

import (
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/errors"
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

// BeforeCreate will check if the username is duplicate.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var count int64
	tx.Model(&User{}).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return errors.WithCode(code.CDuplicateUsername, "%s already exist", u.Username)
	}
	return nil
}

// AfterCreate will automatically generate instanceID based on the ID.
func (u *User) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&u).Update("instance_id", NewInstanceID("uid-", u.ID)).Error
}
