package controller

import (
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

// Feed 视频流
// @Summary 视频流接口，主页的视频流
// @Description 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// @Tags 视频接口
// @Accept application/json
// @Produce application/json
// @Param latest_time query string false "可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间"
// @Param token query string true "用户鉴权token"
// @Success 200 {object} FeedResponse
// @Router /douyin/feed/ [GET]
func Feed(c *gin.Context) {
	var uid int64 = 0
	p := NewProxyFeedVideoList(c)
	token, ok := c.GetQuery("token")
	if ok {
		id := checkToken(token)
		if id != -1 {
			uid = id
		}
	}
	err := p.Do(uid)
	if err != nil {
		p.FeedVideoListError(err.Error())
		return
	}
}

func checkToken(token string) (id int64) {
	if claim, ok := middleware.ParseToken(token); ok {
		// token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return -1
		}
		return claim.UserId
	}
	return -1
}

// Do 视频流推送处理
func (p *ProxyFeedVideoList) Do(id int64) error {
	rawTimestamp := "" //p.Query("latest_time")
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
