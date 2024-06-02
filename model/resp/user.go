package resp

import (
	"time"
)

type UserInfoResp struct {
	CreatedAt     time.Time `json:"created_at"`
	Uid           string    `gorm:"type:varchar(50);comment:uid" json:"uid"`
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Nickname      string    `gorm:"type:varchar(100);comment:昵称" json:"nickname"`
	Email         string    `gorm:"type:varchar(100);comment:邮箱" json:"email"`
	Avatar        string    `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Status        int       `gorm:"type:int;comment:状态" json:"status"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(20);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间" json:"last_login_time"`
}
