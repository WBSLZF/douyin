package dao

import (
	"errors"

	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

func SendMessage(message *model.Message) error {
	if message == nil {
		return errors.New("message空指针")
	}
	return model.DB.Transaction(func(tx *gorm.DB) error {
		//添加聊天数据
		if err := tx.Exec("INSERT INTO messages (chat_key, content, create_time) VALUES (?,?,?)", message.ChatKey, message.Content, message.CreateTime).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func MessageList(chatkey string, Messages *[]model.Message) error {
	if Messages == nil {
		return errors.New("Messages空指针")
	}
	if err := model.DB.Raw("select * from messages where chat_key = ? ORDER BY create_time DESC", chatkey).Scan(Messages).Error; err != nil {
		return err
	}
	return nil
}
