package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Uid     string `gorm:"type:varchar(255);comment:用户uid" json:"content"`
	Content string `gorm:"type:varchar(255);comment:内容" json:"content"`
	IPLoc   string `gorm:"type:varchar(255);comment:ip地址" json:"ip_loc"`
	ISTop   string `gorm:"type:varchar(1);comment:置顶" json:"is_top"`
}

type ArticleAttachment struct {
	AID  int    `gorm:"type:int;comment:文章id" json:"aid"`
	Type string `gorm:"type:varchar(1);comment:类型" json:"type"`
	Url  string `gorm:"type:varchar(255);comment:链接地址" json:"url"`
}
