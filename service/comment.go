package service

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

const (
	timeFormat = "2006-01-02"
)

type Comment struct {
	userId      int64
	videoId     int64
	commentId   int64
	commentText string
	actionType  int64

	comment *model.Comment
}

func CommentAdd(vid int64, uid int64, text string) (comment model.Comment, error error) {
	return Comment{userId: uid, videoId: vid, commentText: text}.CommentAddDo()
}

func CommentDel(vid int64, uid int64, commentId int64) (comment model.Comment, error error) {
	return Comment{userId: uid, videoId: vid, commentId: commentId}.CommentDelDo()
}

func (com Comment) CommentAddDo() (comment model.Comment, error error) {
	if com.videoId == 0 {
		return comment, errors.New("视频消失不见了")
	}
	comment = model.Comment{UserInfoId: com.userId, VideoId: com.videoId, Content: com.commentText, CreateDate: time.Now().Format(timeFormat)}
	err := dao.Comment{}.CommentAdd(&comment)
	if err != nil {
		return comment, errors.New("评论失败")
	}
	return comment, nil
}

func (com Comment) CommentDelDo() (comment model.Comment, error error) {
	if com.videoId == 0 {
		return comment, errors.New("视频消失不见了")
	}
	err := dao.Comment{}.QueryCommentById(com.commentId, &comment)
	if err != nil {
		return comment, errors.New("评论不见了")
	}
	err = dao.Comment{}.CommentDel(com.commentId, com.videoId)
	if err != nil {
		return comment, errors.New("删除评论失败")
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
	err := dao.Comment{}.CommentList(vid, &comments.Comments)
	if err != nil {
		return comments.Comments, err
	}

	return comments.Comments, nil
}
