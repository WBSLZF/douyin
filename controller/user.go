package controller

import (
	"net/http"
	"strconv"

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
	model.Response
	*service.UserLoginData
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

	userLoginData, err := service.UserLogin{}.Register(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response:      model.Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserLoginData: userLoginData,
	})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录功能，判断密码是否正确
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param username query string true "账号"
// @Param password query string true "密码"
// @Success 200 {object} UserLoginResponse
// @Router /douyin/user/login/ [post]
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userLoginData, err := service.UserLogin{}.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response:      model.Response{StatusCode: 0, StatusMsg: "登录成功"},
		UserLoginData: userLoginData,
	})
}

type UserinfoResponse struct {
	model.Response
	User model.UserInfo `json:"user"`
}

// UserInfo 用户信息
// @Summary 获取用户的基本信息
// @Description 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param user_id query string true "用户id"
// @Param token query string true "用户鉴权token"
// @Success 200 {object} UserinfoResponse
// @Router /douyin/user/ [GET]
func UserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserinfoResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "用户id转换异常"},
		})
	}
	ownIdAny, token_user_id_exist := c.Get("user_id")
	if !token_user_id_exist {
		c.JSON(http.StatusOK, UserinfoResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	}
	ownId := ownIdAny.(int64)

	user, err := service.Userinfo{}.SelectUserInfoById(userId, ownId)
	if err != nil {
		c.JSON(http.StatusOK, UserinfoResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	c.JSON(http.StatusOK, UserinfoResponse{
		Response: model.Response{StatusCode: 0},
		User:     user,
	})
}
