package router

import (
	"bobo_server/api/admin"
	"bobo_server/api/front"
	"bobo_server/config"
	"bobo_server/dao"
	"bobo_server/utils"
	"log"
	"net/http"
	"time"
)

var (
	userApi    front.User
	adminApi   admin.Admin
	articleApi front.Article
)

// InitAll 初始化全局变量
func InitAll() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化数据库 DB
	dao.DB = utils.InitMySQLDB()
	// 初始化 Redis
	utils.InitRedis()
}

// FrontServer 前台服务
func FrontServer() *http.Server {
	frontPort := config.Cfg.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", frontPort)
	return &http.Server{
		Addr:         frontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// AdminServer 后台服务
func AdminServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      AdminRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
