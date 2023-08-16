package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response  model.Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

type VideoResponse struct {
	Response  model.Response
	VideoList []model.Video `json:"video_list"`
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")

	videolist, err := service.Videos{}.Getlist(userId, token)
	if err != nil {
		c.JSON(http.StatusOK, VideoResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	c.JSON(http.StatusOK, VideoResponse{
		Response:  model.Response{StatusCode: 1},
		VideoList: videolist,
	})
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})
}
