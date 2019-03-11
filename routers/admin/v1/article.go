package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistArticleByID(id) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, nil)
		return
	}

	data := models.GetArticle(id)
	appG.Response(http.StatusOK, e.Success, data)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	q := c.Query("q")

	var tagId int = -1
	if arg := c.Query("tagId"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	}

	var articleStatus int = -1
	if arg := c.Query("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		maps["article_status"] = articleStatus
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetArticles(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, q, false)
	data["total"] = models.GetArticleTotal(maps, q)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	appG.Response(http.StatusOK, e.Success, data)
}

//新增文章
func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	tagId := com.StrTo(c.PostForm("tagId")).MustInt()
	title := c.PostForm("title")
	sketch := c.PostForm("sketch")
	content := c.PostForm("content")
	weight := com.StrTo(c.DefaultPostForm("weight", "1")).MustInt()
	articleStatus := com.StrTo(c.DefaultPostForm("articleStatus", "1")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(sketch, "sketch").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Range(articleStatus, 0, 100, "weight").Message("权重只允许在0到100之间")
	valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistTagByID(tagId) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, nil)
		return
	}

	data := make(map[string]interface{})
	data["tag_id"] = tagId
	data["title"] = title
	data["sketch"] = sketch
	data["content"] = content
	data["weight"] = weight
	data["article_status"] = articleStatus

	article, ok := models.AddArticle(data)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorArticleCreateFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, *article)
}

//修改文章
func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tagId")).MustInt()
	title := c.PostForm("title")
	sketch := c.PostForm("sketch")
	content := c.PostForm("content")

	var weight int = -1
	if arg := c.PostForm("weight"); arg != "" {
		weight = com.StrTo(arg).MustInt()
		valid.Range(weight, 0, 100, "weight").Message("状态只允许0或100")
	}

	var articleStatus int = -1
	if arg := c.PostForm("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(sketch, 255, "sketch").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistArticleByID(id) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, nil)
		return
	}

	if !models.ExistTagByID(tagId) {
		appG.Response(http.StatusOK, e.ErrorTagNotExists, nil)
		return
	}

	data := make(map[string]interface{})
	if tagId > 0 {
		data["tag_id"] = tagId
	}
	if title != "" {
		data["title"] = title
	}
	if sketch != "" {
		data["sketch"] = sketch
	}
	if weight != -1 {
		data["weight"] = weight
	}
	if content != "" {
		data["content"] = content
	}
	if articleStatus != -1 {
		data["article_status"] = articleStatus
	}

	articleInfo, ok := models.EditArticle(id, data)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorArticleUpdateFailed, data)
		return
	}

	appG.Response(http.StatusOK, e.Success, *articleInfo)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistArticleByID(id) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, nil)
		return
	}

	models.DeleteArticle(id)
	appG.Response(http.StatusOK, e.Success, nil)
}
