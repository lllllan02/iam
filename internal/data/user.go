package data

import (
	"context"

	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/pkg/errors"
)

type UserData interface {
	// Create execute create user.
	Create(c context.Context, user *model.User) error
}

type userData struct{ *Data }

func NewUserData(data *Data) UserData {
	return &userData{Data: data}
}

func (u *userData) Create(c context.Context, user *model.User) error {
	if err := u.DB(c).Create(&user).Error; err != nil {
		return errors.Wrap(err, "create user")
	}
	return nil
}
