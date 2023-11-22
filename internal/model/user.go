package model

import (
	"gorm.io/gorm"
)

type User struct {
	Model

	Nickname string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// AfterCreate run after create database record.
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = NewInstanceID("uid-", u.ID)

	return tx.Save(u).Error
}
