package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	model.Response
}

// FavoriteAction no practical effect, just check if token is valid
// FavoriteAction 用户点赞
// @Summary 用户点赞
// @Description 用户点赞功能，当用户点赞时小红心会点亮，作品点赞数量会+1,点赞列表会多一条记录，取消点赞则反之
// @Tags 互动接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param action_type query string true "点赞标志 1-点赞 2-取消点赞"
// @Success 200 {object} FavoriteActionResponse
// @Router /douyin/favorite/action/ [post]
func FavoriteAction(c *gin.Context) {
	var uid int64
	id, flag := c.Get("user_id")
	if !flag {
		favoriteActionError(c, "用户不存在")
		return
	}
	if id != -1 {
		uid = id.(int64)
		vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
		FavoriteActionDo(c, vid, uid, actionType)
		return
	}
	favoriteActionError(c, "用户不存在")
}

func FavoriteActionDo(c *gin.Context, vid int64, uid int64, actionType int64) {
	err := service.FavoriteAction(vid, uid, actionType)
	if err != nil {
		if actionType == 1 {
			favoriteActionError(c, "点赞失败")
			return
		} else {
			favoriteActionError(c, "取消点赞失败")
			return
		}
	}
	if actionType == 1 {
		favoriteActionOK(c, "点赞成功")
	} else {
		favoriteActionOK(c, "取消点赞成功")
	}

}

func favoriteActionOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FavoriteActionResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  msg,
		},
	})
}

func favoriteActionError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FavoriteActionResponse{
		Response: model.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

// FavoriteList all users have same favorite video list
type FavoriteListResponse struct {
	Response model.Response
	Videos   []*model.Video `json:"video_list,omitempty"`
}

// FavoriteAction 用户喜欢列表
// @Summary 用户喜欢列表
// @Description 查看用户的所有点赞视频
// @Tags 互动接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param user_id query string true "用户id"GET
// @Success 200 {object} FavoriteListResponse
// @Router /douyin/favorite/list/ [GET]
func FavoriteList(c *gin.Context) {
	user_id, exist := c.Get("user_id")
	if !exist {
		favoriteListError(c, "该用户不存在")
		return
	}
	videoList, err := FavoriteListDo(c, user_id.(int64))
	if err != nil {
		favoriteListError(c, "查询用户喜欢列表失败")
		return
	}
	favoriteListOK(c, "查询成功", videoList)
	// token := c.Query("token")

	// if _, exist := usersLoginInfo[token]; exist {
	// 	var uid int64 = 0
	// 	id, err := checkToken(token)
	// 	if err != nil {
	// 		favoriteActionError(c, err.Error())
	// 		return
	// 	}
	// 	if id != -1 {
	// 		uid = id
	// 		videoList, err := FavoriteListDo(c, uid)
	// 		if err != nil {
	// 			favoriteListError(c, "查询失败")
	// 			return
	// 		}
	// 		favoriteListOK(c, "成功查询到喜爱视频列表", videoList)
	// 	}
	// }
	// favoriteListError(c, "用户不存在")
}

func FavoriteListDo(c *gin.Context, uid int64) (videos []*model.Video, error error) {
	videoList, err := service.FavoriteList(uid)
	if err != nil {
		return nil, err
	}
	return videoList.Videos, nil
}

func favoriteListOK(c *gin.Context, msg string, videoList []*model.Video) {
	c.JSON(http.StatusOK, FavoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  msg,
		},
		Videos: videoList,
	},
	)
}

func favoriteListError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FavoriteActionResponse{
		Response: model.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}
