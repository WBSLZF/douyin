package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/utils"
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
	//2.fix 对存储在数据库密码进行加密
	password = utils.MakePassWord(password)
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
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return nil, errors.New("token创建失败")
	}
	userId := userInfo.Id
	//3. 返回需要的数据，对数据进行封装
	return &UserLoginData{Token: token, UserId: userId}, nil
}

func (u UserLogin) Login(name, password string) (*UserLoginData, error) {
	//1. 验证参数是否合法
	if name == "" {
		return nil, errors.New("账号为空")
	}
	if password == "" {
		return nil, errors.New("密码为空")
	}
	//2. 对数据库进行操作
	//2.1 判断用户是否已经存在了
	userExist := dao.UserInfoDao{}.IsUserInfoExistByName(name)
	if !userExist {
		return nil, errors.New("该账号不存在")
	}
	//2.2 判断密码是否相等
	userLogin := dao.UserLoginDao{}.FindUserLoginByName(name)

	//2.fix 加密后再判断是否相等
	if utils.ValidPassWord(password, userLogin.PassWord) {
		return nil, errors.New("密码错误")
	}

	//2.3 创建token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return nil, errors.New("创建token失败")
	}
	//3. 返回前端需要的数据，对数据进行封装
	return &UserLoginData{Token: token, UserId: userLogin.UserInfoId}, nil
}

type UserInfoData struct {
	UserInfo model.UserInfo `json:"user"`
}
