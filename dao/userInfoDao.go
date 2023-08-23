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

// 根据id查找基本信息
func (u UserInfoDao) GetInfoById(userid int64) model.UserInfo {
	userinfo := model.UserInfo{}
	model.DB.Where(userid).Find(&userinfo)

	return userinfo
}

type Relation struct {
	UserInfoId int64 `json:"user_info_id"`
	FollowId   int64 `json:"follow_id"`
}

// 根据两个id判断是否follow
func (u UserInfoDao) IsFollow(own_id, userid int64) bool {
	relation := Relation{own_id, userid}
	result := model.DB.Table("user_relations").Find(relation)
	if result.Error == nil {
		return true
	} else {
		return false
	}
}
