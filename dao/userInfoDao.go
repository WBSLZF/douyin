package dao

import (
	"errors"

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

func (u *UserInfoDao) QueryUserInfoById(userId int64, userinfo *model.UserInfo) error {
	if userinfo == nil {
		return nil
	}
	//DB.Where("id=?",userId).First(userinfo)
	model.DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}
