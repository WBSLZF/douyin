package service

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type Userinfo struct {
}

func (u Userinfo) SelectUserInfoById(userid, ownId int64) (model.UserInfo, error) {

	//根据传入的userid获取基本信息
	user := dao.UserInfoDao{}.GetInfoById(userid)

	//根据自己的ownid和对象的userid判断是否follow
	user.IsFollow = dao.UserInfoDao{}.IsFollow(ownId, userid)
	fmt.Println("-------------------------------------------")
	fmt.Println(user)
	return user, nil
}
