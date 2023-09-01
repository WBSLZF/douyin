package service

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func SendMessage(chatkey string, content string) error {
	var message = model.Message{ChatKey: chatkey, Content: content, CreateTime: time.Now().Format(time.Kitchen)}
	err := dao.SendMessage(&message)
	if err != nil {
		return errors.New("消息发送失败")
	}
	return nil
}

func MessageList(chatkey string) (messageList []*model.Message, error error) {
	messlist, err := dao.MessageList(chatkey)
	if err != nil {
		return nil, err
	}
	return messlist, nil
}
