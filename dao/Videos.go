package dao

import "github.com/RaymondCode/simple-demo/model"

type Vediosdao struct {
}

func (v Vediosdao) GetVideosByUser(user model.UserInfo) []model.Video {

	//获取vedio
	var videos []model.Video
	model.DB.Preload("User").Find(&videos)

	return videos
}
