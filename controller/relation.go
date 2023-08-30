package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	model.Response
	UserList []*model.UserInfo `json:"user_list"`
}

// RelationAction 关注操作
// @Summary 用户与用户之间的关注功能
// @Description 用户用户之间关注
// @Tags 社交接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query string true "对方用户id"
// @Param action_type query string true "1-关注，2-取消关注"
// @Success 200 {object} model.Response
// @Router /douyin/relation/action/ [POST]
func RelationAction(c *gin.Context) {
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type_string := c.Query("action_type")
	action_type, _ := strconv.ParseInt(action_type_string, 10, 64)
	user_id, user_id_exist := c.Get("user_id")

	if !user_id_exist {
		RelationActionErr(c, "用户不存在")
		return
	}
	if user_id == to_user_id {
		RelationActionErr(c, "用户不能关注自己")
		return
	}

	err := service.RelationAction(user_id.(int64), to_user_id, action_type)
	if err != nil {
		RelationActionErr(c, "关注或者取关操作失败")
		return
	}
	RelationActionOK(c, "关注成功")
}

// FollowList 关注列表
// @Summary 查找用户所关注的所有人
// @Description 查找用户所关注的所有人
// @Tags 社交接口
// @Accept application/json
// @Produce application/json
// @Param user_id query string true "用户id"
// @Param token query string true "用户鉴权token"
// @Success 200 {object} UserListResponse
// @Router /douyin/relation/follow/list/ [GET]
func FollowList(c *gin.Context) {
	user_id_string := c.Query("user_id")
	user_id, _ := strconv.ParseInt(user_id_string, 10, 64)

	userList, err := service.Follows{}.FollowList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "查找用户关注列表失败",
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FollowerList 粉丝列表
// @Summary 查找用户的所有粉丝
// @Description 查找用户的所有粉丝
// @Tags 社交接口
// @Accept application/json
// @Produce application/json
// @Param user_id query string true "用户id"
// @Param token query string true "用户鉴权token"
// @Success 200 {object} UserListResponse
// @Router /douyin/relation/follower/list/ [GET]
func FollowerList(c *gin.Context) {
	user_id_string := c.Query("user_id")
	user_id, _ := strconv.ParseInt(user_id_string, 10, 64)

	userList, err := service.Follows{}.FollowerList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "查找用户粉丝列表失败",
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	// c.JSON(http.StatusOK, UserListResponse{
	// 	Response: model.Response{
	// 		StatusCode: 0,
	// 	},
	// 	UserList: []User{DemoUser},
	// })
}
func RelationActionOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  msg,
	})
}
func RelationActionErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}
