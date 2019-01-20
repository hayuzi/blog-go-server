package v1

import (
	"github.com/gin-gonic/gin"
	"blog-go-server/pkg/e"
	"blog-go-server/models"
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"blog-go-server/pkg/logging"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.Success
		} else {
			code = e.ErrorArticleNotExists
		}
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

//获取多个文章
func GetArticles(c *gin.Context) {
}

//新增文章
func AddArticle(c *gin.Context) {
}

//修改文章
func EditArticle(c *gin.Context) {
}

//删除文章
func DeleteArticle(c *gin.Context) {
}
