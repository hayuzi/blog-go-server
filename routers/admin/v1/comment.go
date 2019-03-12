package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取多个评论
func GetComments(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var commentStatus int = -1
	if arg := c.Query("commentStatus"); arg != "" {
		commentStatus = com.StrTo(arg).MustInt()
		maps["comment_status"] = commentStatus
		valid.Range(commentStatus, 1, 2, "commentStatus").Message("状态只允许1或2")
	}

	var articleId int = -1
	if arg := c.Query("articleId"); arg != "" {
		articleId = com.StrTo(arg).MustInt()
		maps["article_id"] = articleId
		valid.Min(articleId, 1, "articleId").Message("文章ID必须大于0")
	}

	var userId int = -1
	if arg := c.Query("userId"); arg != "" {
		userId = com.StrTo(arg).MustInt()
		maps["user_id"] = userId
		valid.Min(userId, 1, "userId").Message("用户ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	data["lists"] = models.GetComments(util.GetQueryOffset(pageNum, pageSize), pageSize, maps)
	data["total"] = models.GetCommentsTotal(maps)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	appG.Response(http.StatusOK, e.Success, data)
}


func DeleteComment(c *gin.Context){
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistCommentByID(id) {
		appG.Response(http.StatusOK, e.ErrorCommentNotExists, nil)
		return
	}

	models.DeleteComment(id)
	appG.Response(http.StatusOK, e.Success, nil)
}
