package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/utils"
)

type VideoPublish struct {
}

// 上传视频，主要是判断上传是否成功
// 根据userid判断用户是否存在，其次用户的filename生成url，自动获取封面，上传到云存储，并添加到数据库中
func (VideoPublish) Upload(userInfoId int64, videoPath string, capturePath string, title string) error {
	//1. 对token进行鉴权，主要是判断用户是否user_id是否存在
	userLogin := dao.UserLoginDao{}.FindUserLoginByUserInfoId(userInfoId)
	if userLogin.UserCount == "" {
		return errors.New("请登录用户")
	}

	// 1.x 准备参数并检验是否合法
	videoUrl, err := utils.GetFileUrl(videoPath)
	if err != nil {
		return err
	}
	captrueUrl, err := utils.GetFileUrl(capturePath)
	if err != nil {
		return err
	}

	//2. 对数据库进行可持久化操作
	//2.1 用户投稿视频添加到video数据库
	err = dao.Video{}.AddVideo(userInfoId, videoUrl, captrueUrl, title)
	if err != nil {
		return err
	}
	//2.2 用户相应的投稿数量进行更新
	userInfo := dao.UserInfoDao{}.GetInfoById(userInfoId)
	userInfo.WorkCount += 1
	err = dao.UserInfoDao{}.UpdateUserInfo(userInfo)
	if err != nil {
		return err
	}
	//3. 返回Controller层需要的数据
	return nil
}
