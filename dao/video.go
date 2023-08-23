package dao

import (
	"time"

	"github.com/RaymondCode/simple-demo/model"
)

type Video struct {
}

func (u Video) AddVideo(userInfoId int64, playUrl string, coverUrl string) error {
	result := model.DB.Create(&model.Video{
		UserInfoId:    userInfoId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateAt:      time.Now(),
	})
	return result.Error
}

func (Video) FindVideoListByUserInfoId(userInfoId int64) (*[]model.Video, error) {
	videoList := []model.Video{}

	result := model.DB.Where("user_info_id = ?", userInfoId).Find(&videoList)
	return &videoList, result.Error
}
