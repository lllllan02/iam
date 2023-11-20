package code

import "github.com/lllllan02/iam/pkg/errors"

var (
	successCoder = Coder{CSuccess, "Success", ""}
	unknowCoder  = Coder{CUnknow, "An internal server error occurred", ""}

	codes = map[int]Coder{
		CSuccess: successCoder,
		CUnknow:  unknowCoder,
		C_ExampleProject_ExampleModule_ExampleErr: {C_ExampleProject_ExampleModule_ExampleErr, "ExampleProject.ExampleModule.ExampleError", ""},
	}
)

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
	return unknowCoder
}
