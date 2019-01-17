package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	TagStatus int    `json:"tagStatus" gorm:"default:1"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now().Unix())
	return nil
}

func GetTags(offset int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(offset).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByTagName(tagName string) bool {
	var tag Tag
	db.Select("id").
		Where("tag_name = ?", tagName).
		First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").
		Where("id = ?", id).
		First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func AddTag(tagName string, TagStatus int) bool {
	db.Create(&Tag{
		TagName:   tagName,
		TagStatus: TagStatus,
	})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
