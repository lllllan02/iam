package repository

import (
	"context"

	"github.com/lllllan02/iam/internal/model"
	"gorm.io/gorm"
)

type MockUserRepo struct{}

func (u MockUserRepo) Create(context.Context, *model.User) error {
	return nil
}

func (u MockUserRepo) First(context.Context, ...func(*gorm.DB) *gorm.DB) (*model.User, error) {
	return &model.User{}, nil
}

func (u MockUserRepo) WithUsername(names ...string) func(*gorm.DB) *gorm.DB {
	return nil
}
