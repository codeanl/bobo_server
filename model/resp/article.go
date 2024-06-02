package resp

import (
	"bobo_server/model"
	"time"
)

type ArticleListResp struct {
	ID                int                        `json:"id"`
	CreatedAt         time.Time                  `json:"created_at"`
	Uid               string                     `json:"uid"`
	Content           string                     `json:"content"`
	IPLoc             string                     `json:"ip_loc"`
	ISTop             string                     `json:"is_top"`
	User              UserInfoResp               `json:"user" gorm:"foreignKey:Uid"`
	ArticleAttachment []*model.ArticleAttachment `json:"article_attachment" gorm:"foreignKey:AID"`
	CommentCount      int                        `json:"comment_count"`    //评论数
	CollectionCount   int                        `json:"collection_count"` //收藏数
	ShareCount        int                        `json:"share_count"`      //分享数
	UpvoteCount       int                        `json:"upvote_count"`     //点赞数
	ViewCount         int                        `json:"view_count"`       //浏览数
}
