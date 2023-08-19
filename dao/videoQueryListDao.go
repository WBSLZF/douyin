package dao

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/model"
)

type VideoDAO struct {
}

// QueryVideoListByLatestTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func (v *VideoDAO) QueryVideoListByLatestTime(limit int, latestTime time.Time, videoList *[]*model.Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}

	return model.DB.Model(&model.Video{}).Where("create_at < ?", latestTime).
		Order("create_at ASC").Limit(limit).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "create_at"}).
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
	if err := model.DB.Raw("SELECT COUNT(*) FROM favor_video WHERE video_id = ? AND user_info_id = ?", videoId, userId).Error; err != nil {
		return true
	}
	return true
}
