package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	adminV1 "blog-go-server/routers/admin/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param id query string false "id"
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
	var id int = -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
		maps["id"] = id
		valid.Min(id, 1, "id").Message("ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	maps["tag_status"] = models.TagStatusNormal

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, tagName, false)
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
