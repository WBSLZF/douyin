package service

import (
	"errors"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

type Comment struct {
	userId      int64
	videoId     int64
	commentId   int64
	commentText string
	actionType  int64

	comment *model.Comment
}

func CommentAction(vid int64, uid int64, commentId int64, text string, actionType int64) (comment model.Comment, error error) {
	return Comment{userId: uid, videoId: vid, commentId: commentId, commentText: text, actionType: actionType}.CommentActionDo()
}

func (com Comment) CommentActionDo() (comment model.Comment, error error) {
	if com.videoId == 0 {
		return comment, errors.New("视频消失不见了")
	}
	if com.actionType == 1 {
		comment := model.Comment{UserInfoId: com.userId, VideoId: com.videoId, Content: com.commentText}
		err := dao.CommentActionY(&comment)
		if err != nil {
			return comment, errors.New("评论失败")
		}
		return comment, nil
	}

	if com.actionType == 2 {
		err := dao.CommentActionN(com.commentId, com.videoId)
		if err != nil {
			return comment, errors.New("删除评论失败")
		}
	}
	return comment, nil
}

type CommList struct {
	videoId int64

	Comments []*model.Comment `json:"video_list"`
}

func CommentList(vid int64) (commentlist []*model.Comment, error error) {
	if vid == 0 {
		return commentlist, errors.New("视频不存在")
	}
	comments := CommList{videoId: vid}
	err := dao.CommentList(vid, &comments.Comments)
	if err != nil {
		return comments.Comments, err
	}
	return comments.Comments, nil
}

// type CommList struct {
// 	Comments []*model.Comment `json:"video_list"`
// }

// type QueryCommentList struct {
// 	userId    int64
// 	videos    []*model.Video
// 	videoList *FavorList
// }

// func CommentList(uid int64) (*CommList, error) {
// 	if uid == 0 {
// 		return nil, errors.New("用户不存在")
// 	}
// 	var q = QueryCommentList{userId: uid}
// 	return q.getCommentList()
// }

// func (q *QueryCommList) getCommentList() (*CommList, error) {
// 	err := dao.CommentList(q.userId, &q.videos)
// 	if err != nil {
// 		return nil, err
// 	}
// 	videodao := dao.VideoDAO{}
// 	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
// 	for i := range q.videos {
// 		//作者信息查询
// 		var userInfo model.UserInfo
// 		err = videodao.QueryUserInfoById(q.videos[i].UserInfoId, &userInfo)
// 		if err == nil { //若查询未出错则更新，否则不更新作者信息
// 			q.videos[i].Author = userInfo
// 		}
// 		q.videos[i].IsFavorite = true
// 	}
// 	q.videoList = &FavorList{Videos: q.videos}
// 	return q.CommList, nil
// }
