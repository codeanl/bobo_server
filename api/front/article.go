package front

import (
	"bobo_server/model/req"
	"bobo_server/utils"
	"bobo_server/utils/r"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// GetArticleList 获取文章列表
func (*Article) GetArticleList(c *gin.Context) {
	data, count := articleService.GetArticleList(utils.BindQuery[req.GetArticleListReq](c))
	r.SuccessListData(c, count, data)
}

// ArticleInfo 文章详情
func (*Article) ArticleInfo(c *gin.Context) {
	data, code := articleService.ArticleInfo(utils.GetIntParam(c, "id"))
	r.SendData(c, code, data)
}

// AddArticle 添加文章
func (*Article) AddArticle(c *gin.Context) {
	code := articleService.AddArticle(c, utils.BindValidJson[req.AddArticleReq](c))
	r.SendCode(c, code)
}
