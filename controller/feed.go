package controller

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response  model.Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
