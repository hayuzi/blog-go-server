package v1

import (
	"blog-go-server/models"
	"blog-go-server/pkg/util"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"net/http"
	serviceCommon "blog-go-server/service/common"
)

type pwdRul struct {
	Pwd      string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
	Confirm  string `valid:"Required; MaxSize(50)"`
}

func ChangePwd(c *gin.Context) {
	appG := app.Gin{C: c}

	pwd := c.PostForm("pwd")           // 旧密码
	password := c.PostForm("password") // 新密码
	confirm := c.PostForm("confirm")   // 新密码确认

	pr := pwdRul{Pwd: pwd, Password: password, Confirm: confirm}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&pr)
	data := make(map[string]interface{})
	code := e.InvalidParams

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}
	if password != confirm {
		appG.Response(http.StatusOK, e.ErrorPasswordDifferent, data)
		return
	}

	claims, _ := serviceCommon.GetLoginClaims(c)
	userInfo := models.GetUser(claims.Id)
	if userInfo.Pwd != util.EncodePwd(pwd) {
		appG.Response(http.StatusOK, e.ErrorUserPassword, data)
		return
	}

	data["pwd"] = util.EncodePwd(password)
	updatedUser, ok := models.EditUser(userInfo.Id, data)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorUserUpdateFailed, data)
		return
	}

	appG.Response(http.StatusOK, code, updatedUser)
}
