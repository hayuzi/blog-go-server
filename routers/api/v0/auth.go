package v0

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Pwd      string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.Query("username")
	pwd := c.Query("pwd")

	valid := validation.Validation{}
	a := auth{Username: username, Pwd: pwd}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	userInfo, isExist := models.CheckAuth(username, pwd)
	if isExist {
		token, err := util.GenerateToken(userInfo.Id, username, pwd)
		if err != nil {
			code = e.ErrorAuthToken
		} else {
			data["token"] = token
			data["id"] = userInfo.Id
			data["username"] = username
			code = e.Success
		}
	} else {
		code = e.ErrorAuth
	}

	appG.Response(http.StatusOK, code, data)
}