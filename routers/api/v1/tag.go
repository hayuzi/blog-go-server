package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	adminV1 "blog-go-server/routers/admin/v1"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param tagName query string false "tagName"
// @Param tagStatus query int false "tagStatus"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id": 2, "createdAt": "2019-01-01 01:16:47", "updatedAt": "2019-01-01 01:16:47", "tagName": "PHP", "weight": 5, "tagStatus": 1}], "pageNum": 1, "pageSize": 10,"total":29},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	tagName := c.Query("tagName")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	if tagName != "" {
		maps["tag_name"] = tagName
	}

	var tagStatus int = -1
	if arg := c.Query("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
		valid.Range(tagStatus, 1, 2, "articleStatus").Message("状态只允许1或2")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), pageSize, maps)
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
	adminV1.GetAllTags(c)
}
