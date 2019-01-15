package models

import "blog-go-server/pkg/constmap"

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	TagStatus int    `json:"tagStatus"`
	delStatus int
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
		Where("del_status = ？", constmap.DelStatusNormal).
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
