package service

import (
	"bobo_server/config"
	"bobo_server/dao"
	"bobo_server/model"
	"bobo_server/model/dto"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type Admin struct{}

// AdminLogin 登录
func (*Admin) AdminLogin(c *gin.Context, username, password string) (token string, code int) {
	// 检查用户是否存在
	user := dao.GetOne(model.Admin{}, "username", username)
	if user.ID == 0 {
		return token, r.ERROR_USER_NOT_EXIST
	}
	if user.Status == 0 {
		return token, r.ERROR_USER_DISABLE
	}
	// 检查密码是否正确
	if password != user.Password {
		return token, r.ERROR_PASSWORD_WRONG
	}
	// 获取 IP 相关信息
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	browser, os := "unknown", "unknown"
	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}
	// 登录信息正确, 生成 Token
	uid := utils.Encryptor.MD5(ipAddress + browser + os) // UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	token, err := utils.GetJWT().GenAdminToken(int(user.ID), uid)
	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return token, r.ERROR_TOKEN_CREATE
	}
	// 更新用户验证信息: ip 信息 + 上次登录时间
	user.IpAddress = ipAddress
	user.IpSource = ipSource
	user.LastLoginTime = time.Now()
	dao.UpdateOne(user, "id = ?", user.ID)
	// 存入redis
	dto := dto.UserRedis{
		ID:            int(user.ID),
		CreatedAt:     user.CreatedAt,
		Uid:           user.Uid,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Email:         user.Email,
		Avatar:        user.Avatar,
		IpAddress:     ipAddress,
		IpSource:      user.IpSource,
		LastLoginTime: user.LastLoginTime,
		Token:         token,
	}
	utils.Redis.Set(KEY_USER+uid, utils.Json.Marshal(dto), time.Duration(config.Cfg.Session.MaxAge)*time.Second)
	//
	return token, r.OK
}

// Profile 个人详情
func (*Admin) Profile(userID int) (user model.Admin, code int) {
	// 检查用户是否存在
	user = dao.GetOne(model.Admin{}, "id", userID)
	if user.ID == 0 {
		return user, r.ERROR_USER_NOT_EXIST
	}
	return user, r.OK
}
