package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/logging"

	"blog-go-server/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"blog-go-server/pkg/e"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param tagName query string false "tagName"
// @Param tagStatus query int false "tagStatus"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id": 2, "createdAt": "2019-01-01 01:16:47", "updatedAt": "2019-01-01 01:16:47", "tagName": "PHP", "weight": 5, "tagStatus": 1}], "pageNum": 1, "pageSize": 10,"total":29},"msg":"ok"}"
// @Router /admin/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}

	tagName := c.Query("tagName")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	valid := validation.Validation{}
	var id int = -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
		maps["id"] = id
		valid.Min(id, 1, "id").Message("ID必须大于0")
	}

	var tagStatus int = -1
	if arg := c.Query("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}
	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, tagName, true)
	data["total"] = models.GetTagTotal(maps)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	appG.Response(http.StatusOK, e.Success, data)
}

// @Summary 获取所有文章标签
// @Produce  json
// @Param tagStatus query int false "tagStatus"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id": 2, "createdAt": "2019-01-01 01:16:47", "updatedAt": "2019-01-01 01:16:47", "tagName": "PHP", "weight": 5, "tagStatus": 1}], "pageNum": 1, "pageSize": 10,"total":29},"msg":"ok"}"
// @Router /admin/v1/tags [get]
func GetAllTags(c *gin.Context) {
	appG := app.Gin{C: c}
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var tagStatus int = -1
	if arg := c.Query("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	data["lists"] = models.GetAllTags(maps)

	appG.Response(http.StatusOK, e.Success, data)
}

// @Summary 新增文章标签
// @Produce  json
// @Param tagName query string true "tagName"
// @Param tagStatus query int false "tagStatus"
// @Param weight query int false "weight"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/tags [post]
func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}

	tagName := c.PostForm("tagName")
	tagStatus := com.StrTo(c.DefaultPostForm("tagStatus", "1")).MustInt()
	weight := com.StrTo(c.DefaultPostForm("weight", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(tagName, "tagName").Message("标签名称不能为空")
	valid.MaxSize(tagName, 60, "tagName").Message("标签名称最长为60字符")
	valid.Range(tagStatus, 1, 2, "tagStatus").Message("状态只允许1或2")
	valid.Range(weight, 0, 100, "tagStatus").Message("权重只允许0到100之间")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if models.ExistTagByTagName(tagName) {
		appG.Response(http.StatusOK, e.ErrorTagExists, nil)
		return
	}

	tagInfo, ok := models.AddTag(tagName, weight, tagStatus)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorTagCreateFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, *tagInfo)
}

// @Summary 修改文章标签
// @Produce  json
// @Param id query int true "ID"
// @Param tagName query string true "id"
// @Param tagName query int false "tagName"
// @Param weight query string true "weight"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	tagName := c.PostForm("tagName")

	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(tagName, 60, "tagName").Message("标签名称最长为60字符")

	var tagStatus int = -1
	if arg := c.PostForm("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		valid.Range(tagStatus, 1, 2, "tagStatus").Message("状态只允许1或2")
	}

	var weight int = -1
	if arg := c.PostForm("weight"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		valid.Range(weight, 0, 100, "tagStatus").Message("权重只允许0到100之间")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistTagByID(id) {
		appG.Response(http.StatusOK, e.ErrorTagNotExists, nil)
		return
	}

	data := make(map[string]interface{})
	if tagStatus != -1 {
		data["tagStatus"] = tagStatus
	}

	_, err := models.EditTag(id, data)
	if err != nil {
		logging.Error(err)
		appG.Response(http.StatusOK, e.ErrorTagUpdateFailed, nil)
		return
	}

	data["id"] = id
	appG.Response(http.StatusOK, e.Success, data)
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if models.ExistTagByID(id) {
		appG.Response(http.StatusOK, e.ErrorTagNotExists, nil)
		return
	}

	models.DeleteTag(id)

	appG.Response(http.StatusOK, e.Success, nil)
}
