package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
)

type Userinfo struct {
}

func (u Userinfo) SelectUserInfoById(userid int64, token string) (model.UserInfo, error) {
	//传入token进行验证
	claims, ok := middleware.ParseToken(token)
	if ok != true {
		return model.UserInfo{}, errors.New("token验证不通过")
	}

	//获取token中的用户自己的id
	own_id := claims.UserId

	//根据传入的userid获取基本信息
	user := dao.UserInfoDao{}.SelectInfoById(userid)

	//根据自己的id判断是否follow
	user.IsFollow = dao.UserInfoDao{}.IsFollow(own_id, userid)

	return user, nil
}
