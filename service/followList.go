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

func (Follows) FriendList(user_id int64) ([]*model.UserInfo, error) {
	result := []*model.UserInfo{}
	userFollowList, err := dao.UserInfoDao{}.FindAllFollow(user_id)
	userFansList, err := dao.UserInfoDao{}.FindAllFollower(user_id)
	idExist := make(map[int64]bool)

	if err != nil {
		return nil, err
	}
	// 查找用户的关注以及粉丝的并集
	for i := range userFollowList {
		follow_id := (*userFollowList[i]).Id
		if _, ok := idExist[follow_id]; !ok {
			(*userFollowList[i]).IsFollow = dao.UserInfoDao{}.IsFollow(user_id, follow_id)
			result = append(result, userFollowList[i])
			idExist[follow_id] = true
		}
	}

	for i := range userFansList {
		fans_id := (*userFansList[i]).Id
		if _, ok := idExist[fans_id]; !ok {
			(*userFansList[i]).IsFollow = dao.UserInfoDao{}.IsFollow(user_id, fans_id)
			result = append(result, userFansList[i])
			idExist[fans_id] = true
		}
	}
	return result, nil
}
