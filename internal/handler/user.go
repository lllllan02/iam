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

// Register godoc
//
//	@Summary		register user
//	@Description	register
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.RegisterReq	true	"register"
//	@Success		200	{object}	resp.Result{data=dto.RegisterRes}
//	@Router			/register [post]
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

// Login godoc
//
//	@Summary		user login
//	@Description	login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.LoginReq	true	"register"
//	@Success		200	{object}	resp.Result{data=dto.LoginRes}
//	@Router			/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	var (
		req *dto.LoginReq
		res *dto.LoginRes
		err error
	)

	if err = c.ShouldBind(&req); err != nil {
		resp.BadRequestError(c)
		return
	}

	res, err = u.userService.Login(c, req)

	resp.JsonResponse(c, err, res)
}
