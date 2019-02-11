package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/logging"
	"blog-go-server/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.Success
		} else {
			data = make(map[string]interface{})
			code = e.ErrorArticleNotExists
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		data = make(map[string]interface{})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var articleStatus int = -1
	if arg := c.Query("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		maps["article_status"] = articleStatus
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	q := c.Query("q")

	var tagId int = -1
	if arg := c.Query("tagId"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	}

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.Success
		pageNum := util.GetPageNum(c)
		pageSize := util.GetPageSize(c)
		data["lists"] = models.GetArticles(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, q)
		data["total"] = models.GetArticleTotal(maps, q)
		data["pageNum"] = pageNum
		data["pageSize"] = pageSize
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
