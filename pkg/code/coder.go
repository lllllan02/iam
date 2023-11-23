package code

import "github.com/lllllan02/iam/pkg/errors"

type Coder struct {
	// 错误码
	Code int

	// 错误描述
	Message string

	// 具体文档地址
	Reference string
}

func ParseCoder(err error) Coder {
	if err == nil {
		return successCoder
	}

	if coder, ok := codes[errors.Code(err)]; ok {
		return coder
	}
	return internalErrCoder
}
