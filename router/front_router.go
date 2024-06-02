package router

import (
	"bobo_server/config"
	"bobo_server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FrontRouter 前台接口路由
func FrontRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		base.POST("/login", userApi.UserLogin)               // 后台登录
		base.GET("/code", userApi.SendCode)                  // 验证码
		base.GET("/article/list", articleApi.GetArticleList) // 文章列表
		base.GET("/article/:id", articleApi.ArticleInfo)     // 文章详情
	}
	// 需要鉴权的接口
	auth := base.Group("") //
	// !注意使用中间件的顺序
	auth.Use(middleware.JWTUserAuth()) // JWT 鉴权中间件
	{
		// 用户模块
		user := auth.Group("/user")
		{
			user.GET("/profile", userApi.Profile)           // 个人详情
			user.POST("/logout", userApi.Logout)            // 退出登录
			user.POST("/up_profile", userApi.UpdateProfile) // 更新个人信息
			user.POST("/set_email", userApi.UpdateEmail)    // 更新个人信息
		}
		//  文章模块
		article := auth.Group("/article")
		{
			article.POST("/add", articleApi.AddArticle) // 添加文章
		}
	}
	return r
}
