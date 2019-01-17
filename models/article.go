package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag

	Title         string `json:"title"`
	Content       string `json:"content"`
	ArticleStatus int    `json:"ArticleStatus"`
	delStatus     int
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now().Unix())
	return nil
}
