package service

import (
	"context"

	"github.com/lllllan02/iam/internal/data"
	"github.com/lllllan02/iam/internal/dto"
	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/pkg/errors"
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
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err = u.userData.Create(c, &user); err != nil {
		return nil, errors.Wrap(err, "UserService.Register")
	}

	res = &dto.RegisterRes{ID: user.ID, UID: user.InstanceID}
	return
}
