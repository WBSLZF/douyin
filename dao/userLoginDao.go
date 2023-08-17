package dao

import "github.com/RaymondCode/simple-demo/model"

type UserLoginDao struct {
}

func (u UserLoginDao) FindUserLoginByName(name string) model.UserLogin {
	userLogin := model.UserLogin{}
	model.DB.Table("user_relation").Where("user_count = ?", name).Find(&userLogin)
	return userLogin
}
