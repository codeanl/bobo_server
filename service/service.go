package service

import "bobo_server/dao"

const (
	KEY_USER               = "user:"              // 记录用户
	KEY_CODE               = "code:"              // 验证码
	KEY_ARTICLE_VIEW_COUNT = "article_view_count" // 文章查看数
)

var (
	articleDao dao.Article
)
