package req

import "bobo_server/model"

type GetArticleListReq struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Uid      string `form:"uid"`
}

type AddArticleReq struct {
	Content           string                    `gorm:"type:varchar(255);comment:内容" json:"content"`
	ArticleAttachment []model.ArticleAttachment `json:"article_attachment"`
}
