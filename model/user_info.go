package model

type UserInfo struct {
	Id            int64       `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	FollowCount   int64       `json:"follow_count,omitempty"`
	FollowerCount int64       `json:"follower_count,omitempty"`
	IsFollow      bool        `json:"is_follow,omitempty"`
	Comments      []*Comment  `json:"-"`                                  //用户有多个评论 一对多
	Videos        []*Video    `json:"-" gorm:"foreignkey:UserInfoId"`     //用户投稿了多个视频 一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"` //用户关注 多对多
	//Messages []*Message `json:"-" gorm:"ma//用户私信 多对多 还没想好该怎么做
	FavorVideos     []*Video   `json:"-" gorm:"many2many:favor_videos;"` //用户点赞视频 多对多
	UserLogin       *UserLogin `json:"-"`                                //用户密码 一对一
	WorkCount       int64      `json:"work_count,omitempty"`             //用户作品数量
	FavoriteCount   int64      `json:"favorite_count,omitempty"`
	TotalFavorited  string     `json:"total_favorited,omitempty"`
	Avatar          string     `json:"avatar,omitempty"`           //用户头像
	BackgroundImage string     `json:"background_image,omitempty"` //用户背景
}
