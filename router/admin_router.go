package router

import (
	"bobo_server/config"
	"bobo_server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminRouter 后台管理接口路由
func AdminRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		base.POST("/login", adminApi.AdminLogin) // 后台登录
	}
	// 需要鉴权的接口
	auth := base.Group("") // "/admin"
	// !注意使用中间件的顺序
	auth.Use(middleware.JWTAdminAuth()) // JWT 鉴权中间件
	{
		// 管理员设置
		admin := auth.Group("/admin")
		{
			admin.GET("/profile", adminApi.Profile) // 个人详情
		}
	}
	return r
}
