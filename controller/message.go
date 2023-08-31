package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

var tempChat = map[string][]model.Message{}

var messageIdSequence = int64(1)

type ChatListResponse struct {
	Response    model.Response
	MessageList []model.Message `json:"message_list"`
}
type ChatResponse struct {
	Response model.Response
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := model.Message{
			Id:         messageIdSequence,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}
		err := service.SendMessage(chatKey, curMessage)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Message send fail"})
			return
		}
		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []model.Message{curMessage}
		}
		c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 0, StatusMsg: chatKey}})
	} else {
		c.JSON(http.StatusOK,
			ChatResponse{Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"}})
	}
}

// 非token测试	token := c.Query("token")
// //userIdA, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
// //userIdB, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
// content := c.Query("content")

// // chatKey := genChatKey(userIdA, userIdB)
// chatKey := "1_3"
// atomic.AddInt64(&messageIdSequence, 1)
// curMessage := model.Message{
// 	Id:         messageIdSequence,
// 	Content:    content,
// 	CreateTime: time.Now().Format(time.Kitchen),
// }
// err := service.SendMessage(chatKey, curMessage)
// if err != nil {
// 	c.JSON(http.StatusOK, ChatResponse{
// 		Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"},
// 	})
// 	return
// }
// if messages, exist := tempChat[chatKey]; exist {
// 	tempChat[chatKey] = append(messages, curMessage)
// } else {
// 	tempChat[chatKey] = []model.Message{curMessage}
// }
// c.JSON(http.StatusOK, ChatResponse{
// 	Response: model.Response{StatusCode: 0, StatusMsg: chatKey},
// })

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {

	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))
		MessageList, err := service.MessageList(chatKey)
		if err == nil {
			tempChat[chatKey] = MessageList
			c.JSON(http.StatusOK, ChatListResponse{
				Response:    model.Response{StatusCode: 0},
				MessageList: tempChat[chatKey]})
		}
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// UserId := c.Query("user_id")
// toUserId := c.Query("to_user_id")
// userIdA, _ := strconv.Atoi(UserId)
// userIdB, _ := strconv.Atoi(toUserId)
// chatKey := genChatKey(int64(userIdA), int64(userIdB))
// MessageList, err := service.MessageList(chatKey)
// if err == nil {
// 	tempChat[chatKey] = MessageList
// 	c.JSON(http.StatusOK, ChatListResponse{
// 		Response:    model.Response{StatusCode: 0},
// 		MessageList: tempChat[chatKey]})
// }

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
