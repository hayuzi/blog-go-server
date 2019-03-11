package models

import (
	"fmt"
)

type Article struct {
	Model
	TagId         int    `json:"tagId" gorm:"index"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title"`
	Sketch        string `json:"sketch"`
	Content       string `json:"content"`
	Weight        int    `json:"weight" gorm:"default:1"`
	ArticleStatus int    `json:"articleStatus" gorm:"default:1"`
}

const (
	ArticleStatusDraft     = 1
	ArticleStatusPublished = 2
)

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").
		Where("id = ?", id).
		First(&article)

	if article.Id > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}, q string) (count int) {
	if q != "" {
		db.Model(&Article{}).Where(maps).Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).Count(&count)
	} else {
		db.Model(&Article{}).Where(maps).Count(&count)
	}
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}, q string, isAdmin bool) (articles []Article) {
	db.Preload("Tag").Where(maps)
	if q != "" {
		db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", q))
	}

	db.Order("weight DESC")
	if isAdmin {
		db.Order("id DESC")
	}
	db.Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).
		First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{}) (*Article, bool) {
	articleInfo := &Article{}
	res := db.Model(articleInfo).Where("id = ?", id).Updates(data)
	if res.RowsAffected == 0 {
		return articleInfo, false
	}
	return articleInfo, true
}

func AddArticle(data map[string]interface{}) (*Article, bool) {
	articleInfo := Article{
		TagId:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Content:       data["content"].(string),
		ArticleStatus: data["article_status"].(int),
	}
	db.Create(&articleInfo)
	return &articleInfo, true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}
