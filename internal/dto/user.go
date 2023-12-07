package dto

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=6,max=60"`
	Email    string `json:"email"`
}

type RegisterRes struct {
	ID  int64  `json:"id"`
	UID string `json:"uid"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=6,max=60"`
}

type LoginRes struct {
	Token string `json:"token"`
}
