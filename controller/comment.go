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

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == 1 {
			text := c.Query("comment_text")
			comment, err := CommentActionAdd(c, videoid, user, text)
			if err == nil {
				commentActionOK(c, comment)
				return
			}
		} else if actionType == 2 {
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			comment, err := CommentActionDel(c, videoid, user, commentId)
			if err == nil {
				commentActionOK(c, comment)
				return
			}
		}
	} else {
		commentActionError(c, "用户解析失败")
	}
}

func CommentActionAdd(c *gin.Context, vid int64, user User, text string) (comment model.Comment, error error) {
	comment, err := service.CommentAdd(vid, user.Id, text)
	if err != nil {
		commentActionError(c, "评论失败")
		return comment, err
	}
	return comment, nil
}

func CommentActionDel(c *gin.Context, vid int64, user User, commentId int64) (comment model.Comment, error error) {
	comment, err := service.CommentDel(vid, user.Id, commentId)
	if err != nil {
		commentActionError(c, "删除评论失败")
		return comment, err
	}
	return comment, nil
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		commentlist, err := CommentListDo(c, videoid)
		if err == nil {
			commentListOK(c, commentlist)
			return
		}
	} else {
		commentListError(c, "用户解析失败")
	}
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
