package service

import (
	context "context"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lllllan02/iam/internal/dto"
	"github.com/lllllan02/iam/internal/service"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/lllllan02/iam/test/data"
	"github.com/stretchr/testify/assert"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	fmt.Println("begin")

	conf := config.NewConfig("../../config/local.yml")
	logger = log.NewLog(conf)

	code := m.Run()

	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTm := data.NewMockTransaction(ctrl)
	mockUserData := data.NewMockUserData(ctrl)

	srv := service.NewService(logger, mockTm)
	userSrv := service.NewUserService(srv, mockUserData)

	ctx := context.Background()
	req := &dto.RegisterReq{
		Username: "test",
		Password: "password",
		Email:    "test@example.com",
	}

	mockTm.EXPECT().Transaction(ctx, gomock.Any()).AnyTimes().Return(nil)

	_, err := userSrv.Register(ctx, req)

	assert.NoError(t, err)
}
