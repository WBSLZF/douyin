package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response  model.Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
// Publish 登录用户选择视频上传
// @Summary 用户投稿
// @Description 投稿首先得鉴权，其次获取用户的上传视频，自动获取封面，上传到云存储，并添加到数据库中
// @Tags 视频
// @Accept multipart/form-data
// @Produce application/json
// @Param data formData file true "视频数据"
// @Param token formData string true "用户鉴权token"
// @Param title formData string true "视频标题"
// @Success 200 {object} model.Response
// @Router /douyin/publish/action/ [post]
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data") //判断表单获取视频是否有问题，需要限制投稿视频的格式要求吗？
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user_id, exist := c.Get("user_id") //判断用户是否登录
	if !exist {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "请登录之后再上传视频",
		})
		return
	}

	// 将视频保存到本地，之后可以考虑上传到云端上面去
	filename := filepath.Base(data.Filename)
	isMP4 := utils.IsMP4File(filename)
	if !isMP4 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "请上传MP4格式的视频",
		})
	}

	finalName := fmt.Sprintf("%d_%s", user_id, filename)
	saveFile := filepath.Join("./public/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 根据视频获取封面，并保存到本地
	saveImgName := utils.ReplaceFileExtension(finalName)
	saveImgPath := filepath.Join("./public/", saveImgName)
	utils.SaveVideoImg(saveFile, saveImgPath)

	// 将保存的文件获取本地访问的url
	err = service.VideoPublish{}.Upload(user_id.(int64), filename, saveImgName, title)
	if err != nil {
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

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
