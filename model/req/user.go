package req

type UserLoginReq struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type UpdateProfileReq struct {
	ID       string `json:"id" label:"id"`
	Nickname string `json:"nickname"  label:"昵称"`
	Avatar   string `json:"avatar" label:"头像"`
}

type UpdateEmailReq struct {
	Email string `json:"email"  label:"邮箱"`
	Code  string `json:"code" label:"验证码"`
}
