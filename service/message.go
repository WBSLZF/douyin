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
	chatkey  string
	Messages []model.Message
}

func MessageList(chatkey string) (messageList []model.Message, error error) {
	messlist := MessList{chatkey: chatkey}
	err := dao.MessageList(chatkey, &messlist.Messages)
	if err != nil {
		return messlist.Messages, err
	}
	return messlist.Messages, nil
}
