package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"blog-go-server/pkg/app"
	"github.com/Unknwon/com"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Pwd      string `valid:"Required; MaxSize(50)"`
}

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

	if username != "" {
		maps["username"] = username
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	pageNum := util.GetPageNum(c)
	pageSize := util.GetPageSize(c)
	users := models.GetUsers(util.GetQueryOffset(pageNum, pageSize), pageSize, maps)
	data["total"] = models.GetUserTotal(maps)
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize

	for key, _ := range users {
		users[key].Pwd = ""
	}

	data["lists"] = users

	appG.Response(http.StatusOK, e.Success, data)
}
