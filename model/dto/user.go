package dto

import "time"

type UserRedis struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Uid           string    `json:"uid"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	Email         string    `json:"email"`
	Avatar        string    `json:"avatar"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	LastLoginTime time.Time `json:"last_login_time"`
	Token         string    `json:"token"`
}
