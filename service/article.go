package service

import (
	"bobo_server/dao"
	"bobo_server/model"
	"bobo_server/model/req"
	"bobo_server/model/resp"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type Article struct{}

// GetArticleList 文章列表
func (*Article) GetArticleList(req req.GetArticleListReq) ([]resp.ArticleListResp, int) {
	list, count := articleDao.GetArticleList(req)
	for item, i := range list {
		list[item].ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(list[item].ID))
		userInfo := dao.GetOne(model.User{}, "uid=?", i.Uid)
		//CopyProperties
		list[item].User = resp.UserInfoResp{
			CreatedAt:     userInfo.CreatedAt,
			Uid:           userInfo.Uid,
			Username:      userInfo.Username,
			Nickname:      userInfo.Nickname,
			Email:         userInfo.Email,
			Avatar:        userInfo.Avatar,
			Status:        userInfo.Status,
			IpAddress:     userInfo.IpAddress,
			IpSource:      userInfo.IpSource,
			LastLoginTime: userInfo.LastLoginTime,
		}
	}
	return list, int(count)
}

// ArticleInfo 文章详情
func (*Article) ArticleInfo(id int) (re resp.ArticleListResp, code int) {
	ar := dao.GetOne(model.Article{}, "id=?", id)
	if ar.ID == 0 {
		return re, r.ERROR_ART_NOT_EXIST
	}
	article := articleDao.ArticleInfo(id)
	// * 目前请求一次就会增加访问量, 即刷新可以刷访问量
	utils.Redis.ZincrBy(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(article.ID), 1)
	userInfo := dao.GetOne(model.User{}, "uid=?", article.Uid)
	copier.Copy(&re, article)
	//CopyProperties
	re.User = resp.UserInfoResp{
		CreatedAt:     userInfo.CreatedAt,
		Uid:           userInfo.Uid,
		Username:      userInfo.Username,
		Nickname:      userInfo.Nickname,
		Email:         userInfo.Email,
		Avatar:        userInfo.Avatar,
		Status:        userInfo.Status,
		IpAddress:     userInfo.IpAddress,
		IpSource:      userInfo.IpSource,
		LastLoginTime: userInfo.LastLoginTime,
	}
	re.ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(id))
	return re, r.OK
}

// AddArticle 添加文章
func (*Article) AddArticle(c *gin.Context, req req.AddArticleReq) (code int) {
	user := dao.GetOne(model.User{}, "id=?", utils.GetFromContext[int](c, "user_info_id"))
	ar := &model.Article{
		Uid:     user.Uid,
		Content: req.Content,
		IPLoc:   utils.IP.GetIpSourceSimpleIdle(utils.IP.GetIpAddress(c)),
		ISTop:   "0",
	}
	dao.Create(ar)
	for item, _ := range req.ArticleAttachment {
		req.ArticleAttachment[item].AID = int(ar.ID)
		dao.Create(&model.ArticleAttachment{
			AID:  req.ArticleAttachment[item].AID,
			Type: req.ArticleAttachment[item].Type,
			Url:  req.ArticleAttachment[item].Url,
		})
	}
	return r.OK
}
