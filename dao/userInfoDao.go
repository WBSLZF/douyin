package dao

import (
	"github.com/RaymondCode/simple-demo/model"
)

type UserInfoDao struct {
}

// 添加用户, 传指针更快一点
func (u UserInfoDao) CreateUserInfo(userinfo *model.UserInfo) error {
	return model.DB.Create(&userinfo).Error
}

func (u UserInfoDao) IsUserInfoExistByName(name string) bool {
	userInfo := model.UserInfo{}
	model.DB.Where("name = ?", name).Find(&userInfo)
	return userInfo.Name != ""
}
