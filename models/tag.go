package models

import (
	"blog-go-server/pkg/e"
	"fmt"
)

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	Weight    int    `json:"weight" gorm:"default:1"`
	TagStatus int    `json:"tagStatus" gorm:"default:1"`
}

const (
	TagStatusNormal = 1
	TagStatusHidden = 2
)

func GetTags(offset int, pageSize int, maps interface{}, q string, isAdmin bool) (tags []Tag) {
	db.Where(maps)
	if q != "" {
		db.Where("tag_name LIKE ?", fmt.Sprintf("%%%s%%", q))
	}
	db.Order("weight DESC")
	if isAdmin {
		db.Order("id DESC")
	}
	db.Offset(offset).Limit(pageSize).Find(&tags)
	return
}

func GetAllTags(maps interface{}) (tags []Tag) {
	db.Where(maps).Order("weight DESC").Order("id DESC").Find(&tags)
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

func AddTag(tagName string, weight int, TagStatus int) (*Tag, bool) {
	tag := Tag{
		TagName:   tagName,
		Weight:    weight,
		TagStatus: TagStatus,
	}
	db.Create(&tag)
	if tag.Id == 0 {
		return &tag, false
	}
	return &tag, true
}

func EditTag(id int, data interface{}) (bool, error) {
	ret := db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	if ret.Error != nil {
		return false, ret.Error
	}
	if ret.RowsAffected == 0 {
		return false, fmt.Errorf("error %d: edit tag failed", e.ErrorTagUpdateFailed)
	}
	return true, nil
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
