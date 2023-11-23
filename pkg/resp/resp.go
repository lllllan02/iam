package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/errors"
)

func JsonResponse(c *gin.Context, err error, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	coder := code.ParseCoder(err)
	c.JSON(200, struct {
		Code      int         `json:"code"`
		Message   string      `json:"message"`
		Reference string      `json:"reference"`
		Data      interface{} `json:"data"`
	}{
		Code:      coder.Code,
		Message:   coder.Message,
		Reference: coder.Reference,
		Data:      data,
	})

	if err != nil {
		c.Error(err)
	}
}

func BadRequestError(c *gin.Context) {
	JsonResponse(c, errors.WithCode(code.CBadRequest, ""), nil)
}
