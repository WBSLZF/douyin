package dao

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

func SendMessage(message model.Message) error {
	return model.DB.Transaction(func(tx *gorm.DB) error {
		//添加聊天数据
		if err := model.DB.Create(&message).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func MessageList(from_user_id, to_user_id int64) (messages []*model.Message, error error) {
	var Messages []*model.Message
	if err := model.DB.Raw("select * from messages where (from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?) ORDER BY time_date DESC", from_user_id, to_user_id, to_user_id, from_user_id).Find(&Messages).Error; err != nil {
		return nil, err
	}
	return Messages, nil
}
