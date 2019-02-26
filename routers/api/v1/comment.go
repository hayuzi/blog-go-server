package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	serviceCommon "blog-go-server/service/common"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个评论
func GetComment(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	var data interface{}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	if !models.ExistCommentByID(id) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, data)
		return
	}

	data = models.GetComment(id)
	appG.Response(http.StatusOK, e.Success, data)
}

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

//新增评论
func AddComment(c *gin.Context) {
	appG := app.Gin{C: c}

	content := c.PostForm("content")
	mentionUserId := com.StrTo(c.DefaultPostForm("mentionUserId", "0")).MustInt()
	articleId := com.StrTo(c.DefaultPostForm("articleId", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(content, "content").Message("评论内容不能为空")
	valid.Min(mentionUserId, 0, "mentionUserId").Message("评论回复用户ID必须大于等于0")
	valid.Min(articleId, 1, "articleId").Message("文章ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if models.ExistArticleByID(articleId) {
		appG.Response(http.StatusOK, e.ErrorArticleNotExists, nil)
		return
	}

	if models.ExistUserByID(mentionUserId) {
		appG.Response(http.StatusOK, e.ErrorUserNotExists, nil)
		return
	}

	claims, _ := serviceCommon.GetLoginClaims(c)
	commentInfo, ok := models.AddComment(articleId, claims.Id, mentionUserId, content)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorCommentCreateFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, *commentInfo)
}
