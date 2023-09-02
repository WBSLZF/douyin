package dao

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type Relations struct {
}

func (Relations) Add(user_id int64, to_user_id int64) error {

	return model.DB.Transaction(func(tx *gorm.DB) error {
		// 插入到关注表中
		if err := tx.Exec("insert into user_relations values (?,?)", user_id, to_user_id).Error; err != nil {
			return err
		}
		// 更新to_user_id的粉丝数
		if err := tx.Exec("update user_infos set follower_count = follower_count + 1 where id = ?", to_user_id).Error; err != nil {
			return err
		}
		// 更新user_id的关注数
		if err := tx.Exec("update user_infos set follow_count = follow_count + 1 where id = ?", user_id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (Relations) Delete(user_id int64, to_user_id int64) error {
	return model.DB.Transaction(func(tx *gorm.DB) error {
		// 根据user_id,to_user_id删除关注表中的数据
		if err := tx.Exec("delete from user_relations where user_info_id = ? and follow_id = ?", user_id, to_user_id).Error; err != nil {
			return err
		}
		// 更新to_user_id的粉丝数
		if err := tx.Exec("update user_infos set follower_count = follower_count - 1 and follower_count >= 0 where id = ?", to_user_id).Error; err != nil {
			return err
		}
		// 更新user_id的关注数
		if err := tx.Exec("update user_infos set follow_count = follow_count - 1 and follow_count >=0 where id = ?", user_id).Error; err != nil {
			return err
		}
		return nil
	})
}
