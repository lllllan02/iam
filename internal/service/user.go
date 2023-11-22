package service

import (
	"context"

	"github.com/lllllan02/iam/internal/data"
	"github.com/lllllan02/iam/internal/dto"
)

type UserService interface {
	// Register is used to register a user.
	Register(c context.Context, req *dto.RegisterReq) (res *dto.RegisterRes, err error)
}

type userService struct {
	*Service

	userData data.UserData
}

func NewUserService(service *Service, userData data.UserData) UserService {
	return &userService{
		Service:  service,
		userData: userData,
	}
}

func (u *userService) Register(c context.Context, req *dto.RegisterReq) (res *dto.RegisterRes, err error) {
	return
}
