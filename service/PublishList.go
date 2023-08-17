package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type Videos struct {
}

func (v Videos) Getvideolist(user model.UserInfo) ([]model.Video, error) {
	//传入token进行验证
	//_, ok := middleware.ParseToken(token)
	//if ok != true {
	//	return []model.Video{}, errors.New("token验证不通过")
	//}

	//根据传入的userid获取基本信息
	//user := dao.UserInfoDao{}.GetInfoById(userid)

	//给dao层获取video
	Videos := dao.Vediosdao{}.GetVideosByUser(user)

	return Videos, nil
}
