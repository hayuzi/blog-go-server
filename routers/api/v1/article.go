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

	var data interface{}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	if !models.ExistArticleByID(id) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, data)
		return
	}

	data = models.GetArticle(id)
	appG.Response(http.StatusOK, e.Success, data)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var tagId int = -1
	if arg := c.Query("tagId"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	}

	q := c.Query("q")
	var articleStatus int = -1
	if arg := c.Query("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		maps["article_status"] = articleStatus
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	maps["articleStatus"] = models.ArticleStatusPublished

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetArticles(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, q, true)
	data["total"] = models.GetArticleTotal(maps, q)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	appG.Response(http.StatusOK, e.Success, data)
}
