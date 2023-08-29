package dao

import (
	"errors"

	"github.com/RaymondCode/simple-demo/model"
)

func FavoriteList(uid int64, videoList *[]*model.Video) error {
	if videoList == nil {
		return errors.New("QueryFavorVideoListByUserId videoList 空指针")
	}
	//多表查询，左连接得到结果，再映射到数据
	if err := model.DB.Raw("SELECT v.* FROM favor_videos u join videos v on u.video_id = v.id WHERE u.user_info_id = ? ", uid).Scan(videoList).Error; err != nil {
		return err
	}
	//如果id为0，则说明没有查到数据
	if len(*videoList) == 0 || (*videoList)[0].Id == 0 {
		return errors.New("点赞列表为空")
	}
	return nil
}
