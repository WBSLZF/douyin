package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	model.Response
	CommentList []*model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction 评论操作
// @Summary 登录用户对视频进行评论
// @Description 已经登录的用户在视频下方进行评论
// @Tags 互动接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param video_id query string true "视频id"
// @Param action_type query string true "1-发布评论 2-删除评论"
// @Param comment_text query string false "用户填写的评论内容，在action_type=1的时候使用"
// @Param comment_id query string false "要删除的评论id，在action_type=2的时候使用"
// @Success 200 {object} CommentActionResponse
// @Router /douyin/comment/action/ [POST]
func CommentAction(c *gin.Context) {
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	user_id, user_id_exist := c.Get("user_id")
	if !user_id_exist {
		commentActionError(c, "用户不存在")
		return
	}
	if actionType == 1 {
		text := c.Query("comment_text")
		comment, err := CommentActionAdd(c, videoid, user_id.(int64), text)
		if err == nil {
			commentActionOK(c, comment)
			return
		}
	} else if actionType == 2 {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		comment, err := CommentActionDel(c, videoid, user_id.(int64), commentId)
		if err == nil {
			commentActionOK(c, comment)
			return
		}
	}
	commentActionError(c, "用户解析失败")
}

func CommentActionAdd(c *gin.Context, vid int64, user_id int64, text string) (comment model.Comment, error error) {
	comment, err := service.CommentAdd(vid, user_id, text)
	if err != nil {
		commentActionError(c, "评论失败")
		return comment, err
	}
	return comment, nil
}

func CommentActionDel(c *gin.Context, vid int64, user_id int64, commentId int64) (comment model.Comment, error error) {
	comment, err := service.CommentDel(vid, user_id, commentId)
	if err != nil {
		commentActionError(c, "删除评论失败")
		return comment, err
	}
	return comment, nil
}

// CommentList 查看用户评论列表
// @Summary 查看视频的所有评论，按发布时间倒序
// @Description 查看视频的所有评论，按发布时间倒序,并不需要限制用户的登录状态吧
// @Tags 互动接口
// @Accept application/json
// @Produce application/json
// @Param token query string true "用户鉴权token"
// @Param video_id query string true "视频id"
// @Success 200 {object} CommentListResponse
// @Router /douyin/comment/list/ [GET]
func CommentList(c *gin.Context) {
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentlist, err := CommentListDo(c, videoid)
	if err == nil {
		commentListOK(c, commentlist)
		return
	}

	commentListError(c, "用户解析失败")
}

// CommentList all videos have same demo comment list
func CommentListDo(c *gin.Context, vid int64) (commentList []*model.Comment, error error) {
	commentList, err := service.CommentList(vid)
	if err != nil {
		commentListError(c, "评论查询失败")
		return commentList, err
	}
	return commentList, nil
}

func commentActionOK(c *gin.Context, comment model.Comment) {
	c.JSON(http.StatusOK, CommentActionResponse{
		Response: model.Response{StatusCode: 0},
		Comment: model.Comment{
			Id:         comment.Id,
			User:       comment.User,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		},
	})
}

func commentActionError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FavoriteActionResponse{
		Response: model.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func commentListOK(c *gin.Context, commentlist []*model.Comment) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: commentlist,
	})
}

func commentListError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FavoriteListResponse{
		Response: model.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}
