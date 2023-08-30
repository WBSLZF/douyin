package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func SendMessage(chatkey string, message model.Message) error {
	message.ChatKey = chatkey
	err := dao.SendMessage(&message)
	if err != nil {
		return errors.New("消息发送失败")
	}
	return nil
}

type MessList struct {
	Messages []model.Message `json:"message_list"`
}

func MessageList(chatkey string) (messageList []model.Message, error error) {
	var Messlist []model.Message
	err := dao.MessageList(chatkey, &Messlist)
	if err != nil {
		return Messlist, err
	}
	return Messlist, nil
}
