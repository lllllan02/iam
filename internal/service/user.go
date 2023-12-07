package service

import (
	"context"
	"time"

	"github.com/lllllan02/iam/internal/dto"
	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/internal/repository"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/lllllan02/iam/pkg/utils/encrypt"
	"github.com/lllllan02/iam/pkg/utils/jwt"
)

func NewUserService(
	conf *config.Config,
	service *Service,
	userRepo repository.UserRepo,
) UserService {
	return &userService{
		Service:  service,
		jwt:      jwt.NewJwt(conf),
		userRepo: userRepo,
	}
}

type userService struct {
	*Service

	jwt      *jwt.JWT
	userRepo repository.UserRepo
}
type UserService interface {
	// Register is used to register a user.
	Register(context.Context, *dto.RegisterReq) (*dto.RegisterRes, error)

	// Login is used for user login
	Login(context.Context, *dto.LoginReq) (*dto.LoginRes, error)
}

func (u *userService) Register(c context.Context, req *dto.RegisterReq) (res *dto.RegisterRes, err error) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// TODO: Send verification code to email

	if err = u.userRepo.Create(c, &user); err != nil {
		return nil, errors.Wrap(err, "register")
	}

	res = &dto.RegisterRes{ID: user.ID, UID: user.InstanceID}
	return
}

func (u *userService) Login(c context.Context, req *dto.LoginReq) (res *dto.LoginRes, err error) {
	user, err := u.userRepo.First(c, u.userRepo.WithUsername(req.Username))
	if err != nil {
		return nil, errors.Wrap(err, "login")
	}

	if user.Password != encrypt.Encrypt(req.Password) {
		return nil, errors.WithCode(code.CIncorrectPassword, "verify password")
	}

	token, err := u.jwt.GenToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, errors.Wrap(err, "generate token")
	}

	res = &dto.LoginRes{Token: token}
	return
}
