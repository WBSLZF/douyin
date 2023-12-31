package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
)

func FavoriteAction(vid int64, uid int64, actionType int64) error {
	if vid == 0 {
		return errors.New("视频消失不见了")
	}
	video, err := dao.Video{}.FindVideoById(vid)
	if err != nil {
		return errors.New("视频不存在")
	}
	ownerId := video.UserInfoId
	if actionType == 1 {
		err := dao.FavoriteActionY(vid, uid, ownerId)
		if err != nil {
			return errors.New("点赞失败")
		}
	}
	if actionType == 2 {
		err := dao.FavoriteActionN(vid, uid, ownerId)
		if err != nil {
			return errors.New("取消点赞失败")
		}
	}
	return nil
}
