package model

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	User       UserInfo `json:"user" gorm:"-;"` // 评论用户作者的相关信息
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`

	UserInfoId int64 `json:"-"` //用户与评论一对多的外键
	VideoId    int64 `json:"-"` //视频与评论一对多的外键
}
