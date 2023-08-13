package model

type UserInfo struct {
	Id            int64       `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	FollowCount   int64       `json:"follow_count,omitempty"`
	FollowerCount int64       `json:"follower_count,omitempty"`
	IsFollow      bool        `json:"is_follow,omitempty"`
	Comments      []*Comment  `json:"-"`                                 //用户有多个评论 一对多
	Videos        []*Video    `json:"-"`                                 //用户投稿了多个视频 一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relation;"` //用户关注 多对多
	//Messages []*Message `json:"-" gorm:"ma//用户私信 多对多 还没想好该怎么做
	FavorVideos []*Video   `json:"-" gorm:"many2many:favor_video;"` //用户点赞视频 多对多
	UserLogin   *UserLogin `json:"-"`                               //用户密码 一对一
}
