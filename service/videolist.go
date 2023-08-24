package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type VideoList struct {
}

func (v VideoList) ListVideo(user_id int64) (*[]model.Video, error) {
	videoList, err := dao.Video{}.FindVideoListByUserInfoId(user_id)
	return videoList, err
}
