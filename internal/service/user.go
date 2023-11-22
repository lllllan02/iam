package service

import "github.com/lllllan02/iam/internal/data"

type UserService struct {
	*Service

	userData *data.UserData
}

func NewUserService(service *Service, userData *data.UserData) *UserService {
	return &UserService{
		Service:  service,
		userData: userData,
	}
}
