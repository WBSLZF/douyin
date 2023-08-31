package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type Follows struct {
}

func (Follows) FollowList(user_id int64) ([]*model.UserInfo, error) {
	userList, err := dao.UserInfoDao{}.FindAllFollow(user_id)
	if err != nil {
		return nil, err
	}
	// 更新下关注信息
	for i := range userList {
		follow_id := (*userList[i]).Id
		(*userList[i]).IsFollow = dao.UserInfoDao{}.IsFollow(user_id, follow_id)
	}
	return userList, nil
}

func (Follows) FollowerList(user_id int64) ([]*model.UserInfo, error) {
	userList, err := dao.UserInfoDao{}.FindAllFollower(user_id)
	if err != nil {
		return nil, err
	}
	for i := range userList {
		fans_id := (*userList[i]).Id
		(*userList[i]).IsFollow = dao.UserInfoDao{}.IsFollow(user_id, fans_id)
	}
	return userList, nil

}
