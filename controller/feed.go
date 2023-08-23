package controller

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

type FeedResponse struct {
	model.Response
	*service.FeedVideoList
}

func Feed(c *gin.Context) {
	var uid int64 = 0
	p := NewProxyFeedVideoList(c)
	token, ok := c.GetQuery("token")
	if ok {
		id, err := checkToken(token)
		if err != nil {
			p.FeedVideoListError(err.Error())
			return
		}
		if id != -1 {
			uid = id
		}
	}
	err := p.FavoriteActionDo(uid)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
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

// Do 视频流推送处理
func (p *ProxyFeedVideoList) FavoriteActionDo(id int64) error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6)
	}
	videoList, err := service.QueryFeedVideoList(id, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

func (p *ProxyFeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{
		Response: model.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyFeedVideoList) FeedVideoListOk(videoList *service.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		FeedVideoList: videoList,
	},
	)
}

type ProxyFeedVideoList struct {
	*gin.Context
}

func NewProxyFeedVideoList(c *gin.Context) *ProxyFeedVideoList {
	return &ProxyFeedVideoList{Context: c}
}
