package dao

import "github.com/RaymondCode/simple-demo/model"

type UserLoginDao struct {
}

func (u UserLoginDao) FindUserLoginByName(name string) model.UserLogin {
	userLogin := model.UserLogin{}
	model.DB.Where("user_count = ?", name).Find(&userLogin)
	return userLogin
}
func (u UserLoginDao) FindUserLoginByUserInfoId(userInfoId int64) model.UserLogin {
	userLogin := model.UserLogin{}
	model.DB.Where("user_info_id = ?", userInfoId).Find(&userLogin)
	return userLogin
}
