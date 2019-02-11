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

	q := c.Query("q")

	var articleStatus int = -1
	if arg := c.Query("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		maps["article_status"] = articleStatus
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

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
		data["total"] = models.GetArticleTotal(maps)
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

//新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tagId")).MustInt()
	title := c.Query("title")
	sketch := c.Query("sketch")
	content := c.Query("content")
	weight := com.StrTo(c.DefaultQuery("weight", "1")).MustInt()
	articleStatus := com.StrTo(c.DefaultQuery("articleStatus", "1")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(sketch, "sketch").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Range(articleStatus, 0, 100, "weight").Message("权重只允许在0到100之间")
	valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["sketch"] = sketch
			data["content"] = content
			data["weight"] = weight
			data["articleStatus"] = articleStatus

			models.AddArticle(data)
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
		"data": make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {

	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tagId")).MustInt()
	title := c.Query("title")
	sketch := c.Query("sketch")
	content := c.Query("content")

	var weight int = -1
	if arg := c.Query("weight"); arg != "" {
		weight = com.StrTo(arg).MustInt()
		valid.Range(weight, 0, 100, "weight").Message("状态只允许0或100")
	}

	var articleStatus int = -1
	if arg := c.Query("articleStatus"); arg != "" {
		articleStatus = com.StrTo(arg).MustInt()
		valid.Range(articleStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(sketch, 255, "sketch").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
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
				if content != "" {
					data["content"] = content
				}
				if articleStatus != -1 {
					data["article_status"] = articleStatus
				}

				models.EditArticle(id, data)
				code = e.Success
			} else {
				code = e.ErrorTagNotExists
			}
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
		"data": make(map[string]interface{}),
	})

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
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
		"data": make(map[string]interface{}),
	})

}
