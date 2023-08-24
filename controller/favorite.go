package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
	})
	// token := c.Query("token")

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
	})
}
