package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	adminV1 "blog-go-server/routers/admin/v1"
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
	adminV1.GetComments(c)
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
