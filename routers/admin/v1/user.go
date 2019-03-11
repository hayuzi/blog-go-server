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

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.Query("username")

	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var userStatus int = -1
	if arg := c.Query("userStatus"); arg != "" {
		userStatus = com.StrTo(arg).MustInt()
		maps["user_status"] = userStatus
		valid.Range(userStatus, 1, 2, "userStatus").Message("状态只允许1或2")
	}

	var userType int = -1
	if arg := c.Query("userType"); arg != "" {
		userType = com.StrTo(arg).MustInt()
		maps["user_type"] = userType
		valid.Min(userType, 1, "userType").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	users := models.GetUsers(util.GetQueryOffset(pageNum, pageSize), pageSize, maps, username)
	data["total"] = models.GetUserTotal(maps, username)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	for k, _ := range users {
		users[k].Pwd = ""
	}

	data["lists"] = users

	appG.Response(http.StatusOK, e.Success, data)
}


func DeleteUser(c *gin.Context){
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	if !models.ExistUserByID(id) {
		appG.Response(http.StatusOK, e.ErrorUserNotExists, nil)
		return
	}

	models.DeleteUser(id)
	appG.Response(http.StatusOK, e.Success, nil)
}
