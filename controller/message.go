package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// var tempChat = map[string][]*model.Message{}

// var messageIdSequence = int64(1)

type ChatListResponse struct {
	model.Response
	MessageList []*model.Message `json:"message_list"`
}
type ChatResponse struct {
	model.Response
}

// MessageAction 发送消息
// @Summary 用户向另外一个用户发送消息
// @Description 用户向另外一个用户发送消息
// @Tags 社交接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query string true "对方用户id"
// @Param action_type query string true "1-发送消息"
// @Param content query string true "消息内容"
// @Success 200 {object} ChatResponse
// @Router /douyin/message/action/ [post]
func MessageAction(c *gin.Context) {
	user_id, user_id_exist := c.Get("user_id")
	if user_id_exist {
		to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
		content := c.Query("content")
		err := service.SendMessage(user_id.(int64), to_user_id, content)
		if err != nil {
			c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"}})
			return
		}
		c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 0, StatusMsg: "Message send success"}})
		return
	}
	c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 1, StatusMsg: "Message send fail"}})

}

// MessageAction 聊天记录
// @Summary 查看与另外一个用户的聊天记录
// @Description 查看与另外一个用户的聊天记录
// @Tags 社交接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query string true "对方用户id"
// @Success 200 {object} ChatListResponse
// @Router /douyin/message/chat/ [get]
func MessageChat(c *gin.Context) {
	user_id, user_id_exist := c.Get("user_id")
	if user_id_exist {
		to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
		pre_msg_time, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64) //前端还有个pre_msg_time 我们可以根据pre_msg_time来保证只添加pre_msg_time之后的消息
		messageList, err := service.MessageList(user_id.(int64), to_user_id, pre_msg_time)
		if err == nil {
			c.JSON(http.StatusOK, ChatListResponse{Response: model.Response{StatusCode: 0},
				MessageList: messageList})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Find Chat Failed"})
		return
	}
	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// func genChatKey(userIdA int64, userIdB int64) string {
// 	if userIdA > userIdB {
// 		return fmt.Sprintf("%d_%d", userIdB, userIdA)
// 	}
// 	return fmt.Sprintf("%d_%d", userIdA, userIdB)
// }
