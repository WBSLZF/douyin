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

// 更新用户
func (u UserInfoDao) UpdateUserInfo(userInfo model.UserInfo) error {
	return model.DB.Save(&userInfo).Error
}

type Relation struct {
	UserInfoId int64 `json:"user_info_id" gorm:"user_info_id"`
	FollowId   int64 `json:"follow_id" gorm:"follow_id"`
}

// 根据两个id判断是否follow
func (u UserInfoDao) IsFollow(own_id, userid int64) bool {
	// relation := Relation{own_id, userid}
	// result := model.DB.Table("user_relations").Find(&relation)
	var n int64
	result := model.DB.Raw("select COUNT(*) from user_relations where user_info_id = ? and follow_id = ?", own_id, userid).Scan(&n)
	if result.Error != nil || n != 0 {
		return true
	} else {
		return false
	}
}

// 根据用户id查找所有其关注的人

func (u UserInfoDao) FindAllFollow(user_id int64) ([]*model.UserInfo, error) {
	var userList []*model.UserInfo
	result := model.DB.Raw("select v.* from user_relations u JOIN user_infos v ON u.follow_id = v.id where u.user_info_id = ?", user_id).Scan(&userList)
	return userList, result.Error
}

func (u UserInfoDao) FindAllFollower(user_id int64) ([]*model.UserInfo, error) {
	var userList []*model.UserInfo
	result := model.DB.Raw("select v.* from user_relations u JOIN user_infos v ON u.user_info_id = v.id where u.follow_id = ?", user_id).Scan(&userList)
	return userList, result.Error
}
