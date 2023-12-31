package dao

import (
	"errors"

	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type Comment struct {
}

func (Comment) CommentAdd(comment *model.Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment空指针")
	}
	return model.DB.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		//增加count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (Comment) CommentDel(commentId, videoId int64) error {
	//执行事务
	return model.DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM `comments` WHERE `id` = ?", commentId).Error; err != nil {
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (Comment) QueryCommentById(id int64, comment *model.Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	return model.DB.Where("id=?", id).First(comment).Error
}

func (Comment) CommentList(videoid int64, Comments *[]*model.Comment) (error error) {
	if Comments == nil {
		return errors.New("QueryCommentListByVideoId comments空指针")
	}
	if err := model.DB.Raw("select * from comments where video_id = ? ORDER BY create_date DESC", videoid).Scan(Comments).Error; err != nil {
		return err
	}
	return nil
}
