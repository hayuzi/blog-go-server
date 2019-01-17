package v1

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"blog-go-server/models"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
)

//获取多个文章标签
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

//新增文章标签
func AddTag(c *gin.Context) {
	tagName := c.PostForm("tagName")
	tagStatus := com.StrTo(c.DefaultPostForm("tagStatus", "1")).MustInt()

	valid := validation.Validation{}
	valid.Required(tagName, "tagName").Message("标签名称不能为空")
	valid.MaxSize(tagName, 60, "tagName").Message("标签名称最长为60字符")
	valid.Range(tagStatus, 1, 2, "tagStatus").Message("状态只允许1或2")

	// Gin 记录日志
	// fmt.Fprintln(gin.DefaultWriter, "foo bar")

	code := e.InvalidParams
	var resData interface{}
	resData = make(map[string]string)

	if !valid.HasErrors() {
		if !models.ExistTagByTagName(tagName) {
			tagInfo, ok := models.AddTag(tagName, tagStatus)
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

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagName := c.PostForm("tagName")

	valid := validation.Validation{}

	var tagStatus int = -1
	if arg := c.PostForm("tagStatus"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		valid.Range(tagStatus, 1, 2, "tagStatus").Message("状态只允许1或2")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(tagName, 60, "tagName").Message("标签名称最长为60字符")

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
