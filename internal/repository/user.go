package repository

import (
	"context"

	"github.com/lllllan02/iam/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	repo *Repo
}

func NewUserRepo(repo *Repo) *UserRepo {
	return &UserRepo{
		repo: repo,
	}
}

func (r UserRepo) First(c context.Context, opts ...Option) (user *model.User, err error) {
	tx := r.repo.DB(c).Model(&model.User{})
	for _, opt := range opts {
		tx = opt(tx)
	}
	err = tx.First(&user).Error
	return
}

func (r UserRepo) WithUsername(names ...string) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("username IN (?)", names)
	}
}
