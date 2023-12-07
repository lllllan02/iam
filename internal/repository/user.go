package repository

import (
	"context"

	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/errors"
	"gorm.io/gorm"
)

func NewUserRepo(data *Repo) UserRepo {
	return &userRepo{Repo: data}
}

type userRepo struct{ *Repo }
type UserRepo interface {
	Create(c context.Context, user *model.User) error
	First(context.Context, ...func(*gorm.DB) *gorm.DB) (*model.User, error)

	// ==================== func(*gorm.DB) *gorm.DB ====================

	WithUsername(names ...string) func(*gorm.DB) *gorm.DB
}

func (u *userRepo) Create(c context.Context, user *model.User) error {
	if err := u.DB(c).Create(&user).Error; err != nil {
		return errors.Wrap(err, "create user")
	}
	return nil
}

func (u *userRepo) First(c context.Context, opts ...func(*gorm.DB) *gorm.DB) (user *model.User, err error) {
	if err = u.DB(c).Scopes(opts...).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.CUsernameNotFound, "first user")
		}
		return nil, errors.Wrap(err, "first user")
	}
	return
}

func (u *userRepo) WithUsername(names ...string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("username IN ?", names)
	}
}
