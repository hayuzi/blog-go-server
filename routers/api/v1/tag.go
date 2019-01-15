package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"

	"blog-go-server/models"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/setting"
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
	if arg := c.Query("tag_status"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
