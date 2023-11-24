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

	if err = u.tm.Transaction(c, func(c context.Context) error {
		if err = u.userData.Create(c, &user); err != nil {
			return errors.Wrap(err, "UserService.Register")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	res = &dto.RegisterRes{ID: user.ID, UID: user.InstanceID}
	return
}
