package model

import "time"

type Video struct {
	Id            int64       `json:"id,omitempty"`
	Author        UserInfo    `json:"author" gorm:"-"` //视频作者的相关信息
	PlayUrl       string      `json:"play_url,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount int64       `json:"favorite_count,omitempty"`
	CommentCount  int64       `json:"comment_count,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty" gorm:"-"`
	Title         string      `json:"title,omitempty"`
	Comments      []*Comment  `json:"-"`                                //视频下有多个评论 一对多
	Users         []*UserInfo `json:"-" gorm:"many2many:favor_videos;"` //用户与视频 多对多
	UserInfoId    int64       `json:"-"`                                //用户投稿多个视频 一对多的外键
	CreateAt      time.Time   `json:"-"`                                //创建时间
}
