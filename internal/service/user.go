package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lllllan02/iam/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(
	userRepo *repository.UserRepo,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s UserService) Login(c context.Context, username, password string) (string, error) {
	user, err := s.userRepo.First(c, s.userRepo.WithUsername(username))
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt)); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["userId"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
