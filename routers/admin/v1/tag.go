package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"fmt"
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
// @Router /admin/v1/tags [get]
func GetTags(c *gin.Context) {
	tagName := c.Query("tagName")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if tagName != "" {
		maps["tag_name"] = tagName
	}

	var tagStatus int = -1
	if arg := c.Query("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	code := e.Success

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), pageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 获取所有文章标签
// @Produce  json
// @Param tagStatus query int false "tagStatus"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id": 2, "createdAt": "2019-01-01 01:16:47", "updatedAt": "2019-01-01 01:16:47", "tagName": "PHP", "weight": 5, "tagStatus": 1}], "pageNum": 1, "pageSize": 10,"total":29},"msg":"ok"}"
// @Router /admin/v1/tags [get]
func GetAllTags(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var tagStatus int = -1
	if arg := c.Query("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	code := e.Success
	data["lists"] = models.GetAllTags(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param tagName query string true "tagName"
// @Param tagStatus query int false "tagStatus"
// @Param weight query int false "weight"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/tags [post]
func AddTag(c *gin.Context) {
	tagName := c.PostForm("tagName")
	tagStatus := com.StrTo(c.DefaultPostForm("tagStatus", "1")).MustInt()
	weight := com.StrTo(c.DefaultPostForm("weight", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(tagName, "tagName").Message("标签名称不能为空")
	valid.MaxSize(tagName, 60, "tagName").Message("标签名称最长为60字符")
	valid.Range(tagStatus, 1, 2, "tagStatus").Message("状态只允许1或2")
	valid.Range(weight, 0, 100, "tagStatus").Message("权重只允许0到100之间")

	// Gin 记录日志
	// fmt.Fprintln(gin.DefaultWriter, "foo bar")

	code := e.InvalidParams
	var resData interface{}
	resData = make(map[string]string)

	if !valid.HasErrors() {
		if !models.ExistTagByTagName(tagName) {
			tagInfo, ok := models.AddTag(tagName, weight, tagStatus)
			if ok {
				code = e.Success
				resData = *tagInfo
			} else {
				code = e.ErrorTagCreateFailed
			}
		} else {
			code = e.ErrorTagExists
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": resData,
	})
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

	fmt.Fprintln(gin.DefaultWriter, tagName)
	fmt.Fprintln(gin.DefaultWriter, c.Param("id"))

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.Success
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			if tagName != "" {
				data["tagName"] = tagName
			}
			if tagStatus != -1 {
				data["tagStatus"] = tagStatus
			}

			fmt.Fprintln(gin.DefaultWriter, tagName)
			fmt.Fprintln(gin.DefaultWriter, c.Param("id"))

			models.EditTag(id, data)
		} else {
			code = e.ErrorTagNotExists
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.Success
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ErrorTagNotExists
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
