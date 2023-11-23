package dto

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email"`
}

type RegisterRes struct {
	ID  int64  `json:"id"`
	UID string `json:"uid"`
}
