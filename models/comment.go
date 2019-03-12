package models

import (
	"blog-go-server/pkg/e"
	"fmt"
)

type Comment struct {
	Model
	ArticleId     int     `json:"articleId" gorm:"index"`
	Article       Article `json:"-"`
	UserId        int     `json:"userId" gorm:"index"`
	User          User    `json:"user"`
	MentionUserId int     `json:"mentionUserId" gorm:"index"`
	MentionUser   User    `json:"mentionUser"gorm:"foreignkey:MentionUserId"`
	Content       string  `json:"content"`
	CommentStatus int     `json:"commentStatus"`
}

const (
	CommentStatusNormal = 1
	CommentStatusHidden = 2
)

func GetComments(offset int, pageSize int, maps interface{}) (comments []Comment) {
	db.Preload("User").
		Preload("MentionUser").
		Where(maps).
		Order("id DESC").
		Offset(offset).Limit(pageSize).Find(&comments)
	return
}

func GetCommentsTotal(maps interface{}) (count int) {
	db.Model(&Comment{}).Where(maps).Count(&count)
	return
}

func ExistCommentByID(id int) bool {
	var comment Comment
	db.Select("id").
		Where("id = ?", id).
		First(&comment)
	if comment.Id > 0 {
		return true
	}
	return false
}

func GetComment(id int) (comment Comment) {
	db.Where("id = ?", id).
		First(&comment)
	db.Model(&comment).Related(&comment.User).Related(&comment.MentionUser)
	return
}

func AddComment(articleId, userId, mentionUserId int, content string) (*Comment, bool) {
	comment := Comment{
		ArticleId:     articleId,
		UserId:        userId,
		MentionUserId: mentionUserId,
		Content:       content,
	}
	db.Create(&comment)
	if comment.Id == 0 {
		return &comment, false
	}
	return &comment, true
}

func EditComment(id int, data interface{}) (bool, error) {
	ret := db.Model(&Comment{}).Where("id = ?", id).Updates(data)
	if ret.Error != nil {
		return false, ret.Error
	}
	if ret.RowsAffected == 0 {
		return false, fmt.Errorf("error %d: edit comment failed", e.ErrorCommentUpdateFailed)
	}
	return true, nil
}

func DeleteComment(id int) bool {
	db.Where("id = ?", id).Delete(Comment{})
	return true
}
