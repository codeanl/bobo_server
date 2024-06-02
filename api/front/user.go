package front

import (
	"bobo_server/model/req"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// UserLogin 登录
func (*User) UserLogin(c *gin.Context) {
	loginReq := utils.BindValidJson[req.UserLoginReq](c)
	loginVo, code := userService.UserLogin(c, loginReq.Username, loginReq.Password)
	r.SendData(c, code, loginVo)
}

// Profile 个人详情
func (*User) Profile(c *gin.Context) {
	profile, code := userService.Profile(utils.GetFromContext[int](c, "user_info_id"))
	r.SendData(c, code, profile)
}

// Logout 退出登录
func (*User) Logout(c *gin.Context) {
	r.SendCode(c, userService.Logout(utils.GetFromContext[string](c, "uuid")))
}

// UpdateProfile 更新个人信息
func (*User) UpdateProfile(c *gin.Context) {
	r.SendCode(c, userService.UpdateProfile(c, utils.BindValidJson[req.UpdateProfileReq](c)))
}

// 发送邮件验证码
func (*User) SendCode(c *gin.Context) {
	r.SendCode(c, userService.SendCode(c.Query("email")))
}

// 发送邮件验证码
func (*User) UpdateEmail(c *gin.Context) {
	r.SendCode(c, userService.UpdateEmail(c, utils.BindValidJson[req.UpdateEmailReq](c)))
}
