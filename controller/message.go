package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// var tempChat = map[string][]*model.Message{}

// var messageIdSequence = int64(1)

type ChatListResponse struct {
	Response    model.Response
	MessageList []*model.Message `json:"message_list"`
}
type ChatResponse struct {
	Response model.Response
}

// MessageAction no practical effect, just check if token is valid ???
func MessageAction(c *gin.Context) {
	user_id, user_id_exist := c.Get("user_id")
	if user_id_exist {
		userIdB, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
		chatKey := genChatKey(user_id.(int64), userIdB)
		content := c.Query("content")
		// atomic.AddInt64(&messageIdSequence, 1)
		// curMessage := model.Message{
		// 	Id:         messageIdSequence,
		// 	Content:    content,
		// 	CreateTime: time.Now().Format(time.Kitchen),
		// }

		err := service.SendMessage(chatKey, content)
		if err != nil {
			c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"}})
		} else {
			c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 0, StatusMsg: "Message send success"}})
		}
		// if messages, exist := tempChat[chatKey]; exist {
		// 	tempChat[chatKey] = append(messages, &curMessage)
		// } else {
		// 	tempChat[chatKey] = []*model.Message{&curMessage}
		// }
	} else {
		c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"}})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	user_id, user_id_exist := c.Get("user_id")
	if user_id_exist {
		userIdB, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
		chatKey := genChatKey(user_id.(int64), userIdB)
		messageList, err := service.MessageList(chatKey)
		if err == nil {
			// tempChat[chatKey] = messageList
			c.JSON(http.StatusOK, ChatListResponse{Response: model.Response{StatusCode: 0},
				MessageList: messageList})
		}
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
