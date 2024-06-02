package model

import (
	"gorm.io/gorm"
)

type Common struct {
	gorm.Model
	Uid     int    `gorm:"type:int;comment:父级id" json:"uid"`
	Aid     int    `gorm:"type:int;comment:文章id" json:"aid"`
	Content string `gorm:"type:varchar(200);comment:内容" json:"content"`
	Pid     int    `gorm:"type:int;comment:父级id" json:"pid"`
	Pic     string `gorm:"type:varchar(255);comment:图片" json:"pic"`
}
