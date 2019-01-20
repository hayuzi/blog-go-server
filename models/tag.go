package models

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	Weight    int    `json:"weight" gorm:"default:1"`
	TagStatus int    `json:"tagStatus" gorm:"default:1"`
}

func GetTags(offset int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Order("weight DESC").Order("id DESC").Offset(offset).Limit(pageSize).Find(&tags)
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
	return &tag, true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
