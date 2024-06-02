package dao

import (
	"bobo_server/model/req"
	"bobo_server/model/resp"
)

type Article struct{}

func (*Article) GetArticleList(req req.GetArticleListReq) ([]resp.ArticleListResp, int64) {
	list := make([]resp.ArticleListResp, 0)
	var total int64

	db := DB.Table("article").Select("*")

	if req.Uid != "" {
		db = db.Where("uid = ?", req.Uid)
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	db.Count(&total).Preload("ArticleAttachment").Order("created_at DESC").Find(&list)
	return list, total
}

func (*Article) ArticleInfo(id int) (resp resp.ArticleListResp) {
	err := DB.Table("article").
		Select("*").
		Where("id = ?", id).
		Preload("ArticleAttachment").
		Order("created_at DESC").
		First(&resp).Error

	if err != nil {
	}
	return resp
}
