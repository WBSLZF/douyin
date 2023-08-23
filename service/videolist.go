package service

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

// MaxVideoNum 每次最多返回的视频流数量
const (
	MaxVideoNum = 30
)

type FeedVideoList struct {
	Videos   []*model.Video `json:"video_list,omitempty"`
	NextTime int64          `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}

type QueryFeedVideoListFlow struct {
	userId     int64
	latestTime time.Time

	videos   []*model.Video
	nextTime int64

	feedVideo *FeedVideoList
}

func NewQueryFeedVideoListFlow(userId int64, latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{userId: userId, latestTime: latestTime}
}

func (q *QueryFeedVideoListFlow) Do() (*FeedVideoList, error) {
	if q.latestTime.IsZero() {
		q.latestTime = time.Now()
	}

	if err := q.prepareData(); err != nil {
		return nil, err
	}
	q.feedVideo = &FeedVideoList{
		Videos:   q.videos,
		NextTime: q.nextTime,
	}
	return q.feedVideo, nil
}

func (q *QueryFeedVideoListFlow) prepareData() error {
	err := dao.NewVideoDAO().QueryVideoListByLatestTime(MaxVideoNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	if q.userId != 1 {
		//如果用户为登录状态，则更新该视频是否被该用户点赞的状态
		latestTime, _ := FillVideoListFields(q.userId, &q.videos)
		//准备好时间戳
		if latestTime != nil {
			q.nextTime = (*latestTime).UnixNano() / 1e6
			return nil
		}
	}
	q.nextTime = time.Now().Unix() / 1e6
	return nil
}

// FillVideoListFields 填充每个视频的作者信息（因为作者与视频的一对多关系，数据库中存下的是作者的id
// 当userId>0时，我们判断当前为登录状态，其余情况为未登录状态，则不需要填充IsFavorite字段
func FillVideoListFields(userId int64, videos *[]*model.Video) (*time.Time, error) {

	if videos == nil {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	size := len(*videos)
	if size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	userdao := dao.UserInfoDao{}
	videodao := dao.VideoDAO{}

	latestTime := (*videos)[size-1].CreateAt //获取最近的投稿时间
	//添加作者信息，以及is_follow状态
	for i := 0; i < size; i++ {
		var userInfo model.UserInfo
		err := userdao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		userInfo.IsFollow = videodao.GetUserRelation(userId, userInfo.Id) //根据cache更新是否被点赞
		(*videos)[i].Author = userInfo
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = videodao.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}
