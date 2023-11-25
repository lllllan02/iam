package code

// 通用错误码
const (
	CSuccess      = 200
	CBadRequest   = 400
	CUnauthorized = 401
	CForbidden    = 403
	CNotFound     = 404
	CInternalErr  = 500

	C_ExampleProject_ExampleModule_ExampleErr = 999999 // 项目编号 99 | 模块编号 99 | 错误编号 99
)

// iam 项目
const (
	// user 模块
	CInvalidUsername   = 10101
	CDuplicateUsername = 10102
	CInvalidEmail      = 10103
	CDuplicaEmail      = 10104
)

var (
	successCoder     = Coder{CSuccess, "Success!", ""}
	internalErrCoder = Coder{CInternalErr, "An internal server error occurred.", ""}
)

var codes = map[int]Coder{
	CSuccess:      successCoder,
	CBadRequest:   {CBadRequest, "[Bad Request] Invalid request parameter.", ""},
	CUnauthorized: {CUnauthorized, "[Unauthorized] Please log in.", ""},
	CForbidden:    {CForbidden, "[Forbidden] Insufficient permissions to access.", ""},
	CNotFound:     {CNotFound, "[Not Found] Record not found.", ""},
	CInternalErr:  internalErrCoder,
	C_ExampleProject_ExampleModule_ExampleErr: {C_ExampleProject_ExampleModule_ExampleErr, "ExampleProject.ExampleModule.ExampleErr", ""},

	CInvalidUsername:   {CInvalidUsername, "Invalid username.", ""},
	CDuplicateUsername: {CDuplicateUsername, "Username already exists!", ""},
	CInvalidEmail:      {CInvalidEmail, "Invalid email", ""},
	CDuplicaEmail:      {CDuplicaEmail, "Email already exists!", ""},
}
