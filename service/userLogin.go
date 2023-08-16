package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
)

type UserLogin struct {
}

type UserLoginData struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func (u UserLogin) Register(name, password string) (*UserLoginData, error) {
	//1. 验证参数是否合法
	if name == "" {
		return nil, errors.New("账号为空")
	}
	if password == "" {
		return nil, errors.New("密码为空")
	}
	//2. 对数据库进行操作，以及获取相应的数据
	userLogin := model.UserLogin{UserCount: name, PassWord: password}
	userInfo := model.UserInfo{UserLogin: &userLogin, Name: name}
	//2.1 判断用户是否已经存在了
	userExist := dao.UserInfoDao{}.IsUserInfoExistByName(name)
	if userExist {
		return nil, errors.New("该账号已经存在了")
	}

	err := dao.UserInfoDao{}.CreateUserInfo(&userInfo)
	if err != nil {
		return nil, err
	}
	//2.2 token创建失败
	token, err := middleware.ReleaseToken(userLogin.UserInfoId)
	if err != nil {
		return nil, errors.New("token创建失败")
	}
	userId := userInfo.Id
	//3. 返回需要的数据，对数据进行封装
	return &UserLoginData{Token: token, UserId: userId}, nil
}

type UserInfoData struct {
	UserInfo model.UserInfo `json:"user"`
}

func (u UserLogin) Login(name, password string) *UserInfoData {
	user := &UserInfoData{}
	//1. 验证参数是否合法
	//2. 对数据库进行操作
	//3. 返回前端需要的数据，对数据进行封装
	return user
}
