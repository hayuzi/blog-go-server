package v1

import (
	"blog-go-server/pkg/constmap"
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
	tagName := c.Query("tag_name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["del_status"] = constmap.DelStatusNormal

	if tagName != "" {
		maps["tag_name"] = tagName
	}

	var tagStatus int = -1
	if arg := c.Query("tag_status"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	code := e.Success

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), pageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	data["page_num"] = pageNum
	data["page_size"] = pageSize

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	tagName := c.Query("tag_name")
	tagStatus := com.StrTo(c.DefaultQuery("tag_status", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(tagName, "tag_name").Message("标签名称不能为空")
	valid.MaxSize(tagName, 60, "tag_name").Message("标签名称最长为60字符")
	valid.Range(tagStatus, 1, 2, "tag_status").Message("状态只允许1或2")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if !models.ExistTagByTagName(tagName) {
			code = e.Success
			models.AddTag(tagName, tagStatus)
		} else {
			code = e.ErrorTagExists
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
