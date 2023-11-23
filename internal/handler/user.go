package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/internal/dto"
	"github.com/lllllan02/iam/internal/service"
	"github.com/lllllan02/iam/pkg/resp"
)

type UserHandler struct {
	*Handler

	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (u *UserHandler) Register(c *gin.Context) {
	var (
		req *dto.RegisterReq
		res *dto.RegisterRes
		err error
	)

	if err = c.ShouldBind(&req); err != nil {
		resp.BadRequestError(c)
		return
	}

	res, err = u.userService.Register(c, req)

	resp.JsonResponse(c, err, res)
}
