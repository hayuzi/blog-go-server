package v1

import (
	"blog-go-server/pkg/constmap"
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
	tagName := c.Query("tag_name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["del_status"] = constmap.DEL_STATUS_NORMAL

	if tagName != "" {
		maps["tag_name"] = tagName
	}

	var tagStatus int = -1
	if arg := c.Query("tag_status"); arg != "" {
		tagStatus = com.StrTo(arg).MustInt()
		maps["tag_status"] = tagStatus
	}

	code := e.SUCCESS

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetTags(util.GetQueryOffset(pageNum, pageSize), setting.PageSize, maps)
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

}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
