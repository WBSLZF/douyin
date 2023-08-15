package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response      model.Response
	UserLoginData *service.UserLoginData
	// UserId int64  `json:"user_id,omitempty"`
	// Token  string `json:"token"`
}

type UserResponse struct {
	Response model.Response
	User     User `json:"user"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册，需要判断用户名是否已经被注册了，以及用户密码是否规范
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param username query string true "账号"
// @Param password query string true "密码"
// @Success 200 {object} UserLoginResponse
// @Router /douyin/user/register/ [post]
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password

	userLoginData, err := service.UserLogin{}.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response:      model.Response{StatusCode: 0, StatusMsg: "注册成功"},
			UserLoginData: userLoginData,
		})
	}
}

func Login(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")

	// token := username + password

	// // if user, exist := usersLoginInfo[token]; exist {
	// // 	c.JSON(http.StatusOK, UserLoginResponse{
	// // 		Response:      model.Response{StatusCode: 0},
	// // 		UserLoginData: &service.UserLogin{},
	// // 	})
	// // } else {
	// // 	c.JSON(http.StatusOK, UserLoginResponse{
	// // 		Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// // 	})
	// // }
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
