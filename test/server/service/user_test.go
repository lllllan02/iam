package service

import (
	"context"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/lllllan02/iam/internal/dto"
	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/internal/service"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/lllllan02/iam/pkg/utils/encrypt"
	"github.com/lllllan02/iam/test/mocks/repository"
)

var (
	srv  = service.NewService(&log.Logger{}, &repository.MockRepo{})
	conf = config.NewConfig("../../../config/local.yml")
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	userSrv := service.NewUserService(conf, srv, &repository.MockUserRepo{})

	// 用户名不存在
	PatchConvey("1", t, func() {
		Mock(repository.MockUserRepo.First).Return(nil, errors.WithCode(code.CUsernameNotFound, "first user")).Build()
		_, err := userSrv.Login(ctx, &dto.LoginReq{Username: "test", Password: "password"})
		So(errors.Code(err), ShouldEqual, code.CUsernameNotFound)
	})

	// 密码错误
	PatchConvey("2", t, func() {
		Mock(repository.MockUserRepo.First).Return(&model.User{Username: "test", Password: "password"}, nil).Build()
		_, err := userSrv.Login(ctx, &dto.LoginReq{Username: "test", Password: "password"})
		So(errors.Code(err), ShouldEqual, code.CIncorrectPassword)
	})

	// 登录成功
	PatchConvey("3", t, func() {
		Mock(repository.MockUserRepo.First).Return(&model.User{Username: "test", Password: encrypt.Encrypt("password")}, nil).Build()
		_, err := userSrv.Login(ctx, &dto.LoginReq{Username: "test", Password: "password"})
		So(err, ShouldEqual, nil)
	})
}
