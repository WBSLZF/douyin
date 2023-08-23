package favorite

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FavoriteListResponse struct {
	Response model.Response
	Videos   []*model.Video `json:"video_list,omitempty"`
}

func favoriteList(c *gin.Context) {
	token, ok := c.GetQuery("token")
	uid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if ok {
		id, err := checkToken(token)
		if err != nil {
			favoriteListError(c, err.Error())
			return
		}
		if id != uid {
			favoriteListError(c, err.Error())
			return
		}
	} else {
		favoriteListError(c, "用户未登录")
		return
	}
	videoList, err := FavoriteListDo(c, uid)
	if err != nil {
		favoriteListError(c, "查询失败")
		return
	}
	favoriteListOK(c, "成功查询到喜爱视频列表", videoList)
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
