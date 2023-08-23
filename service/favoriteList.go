package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type FavorList struct {
	Videos []*model.Video `json:"video_list"`
}

type QueryFavorList struct {
	userId    int64
	videos    []*model.Video
	videoList *FavorList
}

func FavoriteList(uid int64) (*FavorList, error) {
	if uid == 0 {
		return nil, errors.New("用户不存在")
	}
	var q = QueryFavorList{userId: uid}
	return q.getFavorList()
}

func (q *QueryFavorList) getFavorList() (*FavorList, error) {
	err := dao.FavoriteList(q.userId, &q.videos)
	if err != nil {
		return nil, err
	}
	videodao := dao.VideoDAO{}
	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
	for i := range q.videos {
		//作者信息查询
		var userInfo model.UserInfo
		err = videodao.QueryUserInfoById(q.videos[i].UserInfoId, &userInfo)
		if err == nil { //若查询未出错则更新，否则不更新作者信息
			q.videos[i].Author = userInfo
		}
		q.videos[i].IsFavorite = true
	}
	q.videoList = &FavorList{Videos: q.videos}
	return q.videoList, nil
}
