package service

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func SendMessage(user_id int64, to_user_id int64, content string) error {
	var message = model.Message{FromUserId: user_id, ToUserId: to_user_id, Content: content, TimeDate: time.Now().Format(time.DateTime)}
	err := dao.SendMessage(message)
	if err != nil {
		return errors.New("消息发送失败")
	}
	return nil
}

func MessageList(user_id int64, to_user_id int64, pre_msg_time int64) (messageList []*model.Message, error error) {
	messlist, err := dao.MessageList(user_id, to_user_id)
	result := []*model.Message{}
	//返回的是时间戳
	for id := range messlist {
		createTime := Time2Unix((*messlist[id]).TimeDate)
		if createTime > pre_msg_time {
			(*messlist[id]).CreateTime = createTime
			result = append(result, messlist[id])
		}
		// fmt.Println("------------------------------------------------------------")
		// fmt.Println("时间戳是:", (*messlist[id]).CreateTime)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Time2Unix(datetime string) int64 {
	//日期转化为时间戳
	timeLayout := time.DateTime          //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区

	//调用转化方法，传入上面准备好的的三个参数
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳（秒级） 类型是int64
	return timestamp
}
