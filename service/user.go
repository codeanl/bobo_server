package service

import (
	"bobo_server/config"
	"bobo_server/dao"
	"bobo_server/model"
	"bobo_server/model/dto"
	"bobo_server/model/req"
	"bobo_server/model/resp"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type User struct{}

// UserLogin 登录
func (*User) UserLogin(c *gin.Context, username, password string) (token string, code int) {
	// 检查用户是否存在
	user := dao.GetOne(model.User{}, "username", username)
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
func (*User) Profile(userID int) (user resp.UserInfoResp, code int) {
	// 检查用户是否存在
	userinfo := dao.GetOne(model.User{}, "id", userID)
	if userinfo.ID == 0 {
		return user, r.ERROR_USER_NOT_EXIST
	}
	user = resp.UserInfoResp{
		CreatedAt:     userinfo.CreatedAt,
		Uid:           userinfo.Uid,
		Username:      userinfo.Username,
		Nickname:      userinfo.Nickname,
		Email:         userinfo.Email,
		Avatar:        userinfo.Avatar,
		Status:        userinfo.Status,
		IpAddress:     userinfo.IpAddress,
		IpSource:      userinfo.IpSource,
		LastLoginTime: userinfo.LastLoginTime,
	}
	return user, r.OK
}

// Logout 退出登录
func (*User) Logout(uuid string) (code int) {
	utils.Redis.Del(KEY_USER + uuid)
	return r.OK
}

// UpdateProfile 更新个人信息
func (*User) UpdateProfile(c *gin.Context, req req.UpdateProfileReq) (code int) {
	user := utils.CopyProperties[model.User](req)
	dao.UpdateOne(user, "id = ?", utils.GetFromContext[int](c, "user_info_id"))
	return r.OK
}

// 发送验证码
func (*User) SendCode(email string) (code int) {
	// 已经发送验证码且未过期
	if utils.Redis.GetVal(KEY_CODE+email) != "" {
		return r.ERROR_EMAIL_HAS_SEND
	}

	expireTime := config.Cfg.Captcha.ExpireTime
	validateCode := utils.Encryptor.ValidateCode()
	content := fmt.Sprintf(`
		<div style="text-align:center"> 
			<div>你好！欢迎访问阵、雨的个人博客！</div>
			<div style="padding: 8px 40px 8px 50px;">
            	<p>
					您本次的验证码为
					<p style="font-size:75px;font-weight:blod;"> %s </p>
					为了保证账号安全，验证码有效期为 %d 分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用~
				</p>
       	 	</div>
			<div>
            	<p>发送邮件专用邮箱，请勿回复。</p>
        	</div>
		</div>
	`, validateCode, expireTime)

	if err := utils.Email(email, "博客注册验证码", content); err != nil {
		return r.ERROR_EMAIL_SEND
	}

	// 将验证码存储到 Redis 中
	utils.Redis.Set(KEY_CODE+email, validateCode, time.Duration(expireTime)*time.Minute)
	return r.OK
}

// UpdateEmail 更新邮箱
func (*User) UpdateEmail(c *gin.Context, req req.UpdateEmailReq) (code int) {
	// 已经发送验证码且未过期
	if req.Code != utils.Redis.GetVal(KEY_CODE+req.Email) {
		return r.ERROR_VERIFICATION_CODE
	}
	dao.UpdateOne(model.User{Email: req.Email}, "id = ?", utils.GetFromContext[int](c, "user_info_id"))
	return r.OK
}
