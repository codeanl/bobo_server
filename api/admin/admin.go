package admin

import (
	"bobo_server/model/req"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"github.com/gin-gonic/gin"
)

type Admin struct{}

// AdminLogin 登录
func (*Admin) AdminLogin(c *gin.Context) {
	loginReq := utils.BindValidJson[req.AdminLoginReq](c)
	loginVo, code := adminService.AdminLogin(c, loginReq.Username, loginReq.Password)
	r.SendData(c, code, loginVo)
}

// Profile 个人详情
func (*Admin) Profile(c *gin.Context) {
	profile, code := adminService.Profile(utils.GetFromContext[int](c, "user_info_id"))
	r.SendData(c, code, profile)
}
