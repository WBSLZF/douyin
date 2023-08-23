package favorite

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	Response model.Response
}

func favoriteaction(c *gin.Context) {
	token, ok := c.GetQuery("token")
	var uid int64 = 0
	if ok {
		id, err := checkToken(token)
		if err != nil {
			favoriteActionError(c, err.Error())
			return
		}
		if id != -1 {
			uid = id
		}
	} else {
		favoriteActionError(c, "用户未登录")
		return
	}
	vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	Do(c, vid, uid, actionType)
}

func checkToken(token string) (id int64, error error) {
	if claim, ok := middleware.ParseToken(token); ok {
		// token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return -1, errors.New("token超时")
		}
		return claim.UserId, nil
	}
	return -1, errors.New("token不正确")
}

func Do(c *gin.Context, vid int64, uid int64, actionType int64) {
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
		favoriteActionError(c, "点赞成功")
	} else {
		favoriteActionError(c, "取消点赞成功")
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
