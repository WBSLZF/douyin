package dao

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

func FavoriteActionY(vid int64, uid int64, ownerId int64) error {

	return model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count+1 WHERE id = ?", vid).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET favorite_count=favorite_count+1 WHERE id = ?", uid).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `favor_videos` (`user_info_id`,`video_id`) VALUES (?,?)", uid, vid).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET total_favorited=total_favorited+1 WHERE id = ?", ownerId).Error; err != nil {
			return err
		}
		return nil
	})
}

func FavoriteActionN(vid int64, uid int64, ownerId int64) error {
	return model.DB.Transaction(func(tx *gorm.DB) error {
		//执行-1之前需要先判断是否合法（不能被减少为负数
		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count-1 WHERE id = ? AND favorite_count>=0", vid).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET favorite_count=favorite_count-1 WHERE id = ?  AND favorite_count>=0", uid).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `favor_videos`  WHERE `user_info_id` = ? AND `video_id` = ?", uid, vid).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET total_favorited=total_favorited-1 AND total_favorited >=0  WHERE id = ?", ownerId).Error; err != nil {
			return err
		}
		return nil
	})
}
