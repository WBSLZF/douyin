package dao

import (
	"errors"

	// "sync"
	"time"

	"github.com/RaymondCode/simple-demo/model"
)

type VideoDAO struct {
}

func (v *VideoDAO) QueryUserInfoById(userId int64, userinfo *model.UserInfo) error {
	if userinfo == nil {
		return nil
	}
	//DB.Where("id=?",userId).First(userinfo)
	model.DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// QueryVideoListByLatestTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func (v *VideoDAO) QueryVideoListByLatestTime(limit int, latestTime time.Time, videoList *[]*model.Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}
	return model.DB.Model(&model.Video{}).Where("create_at < ?", latestTime).
		Order("create_at ASC").Limit(limit).
		Find(videoList).Error
}

func (v VideoDAO) GetUserRelation(userId int64, userInfoId int64) bool {
	if userId == 0 || userInfoId == 0 {
		return false
	}
	if err := model.DB.Raw("SELECT COUNT(*) FROM user_relation WHERE follow_id = ? AND user_info_id = ?", userId, userInfoId).Error; err != nil {
		return true
	}
	return false
}

func (v VideoDAO) GetVideoFavorState(userId int64, videoId int64) bool {
	if userId == 0 || videoId == 0 {
		return false
	}
	var n int
	if model.DB.Raw("SELECT Count(*) FROM favor_videos WHERE video_id = ? AND user_info_id = ?", videoId, userId).Scan(&n); n == 0 {
		return false
	}
	return true
}
