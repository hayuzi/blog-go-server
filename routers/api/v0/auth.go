package v0

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Pwd      string `valid:"Required; MaxSize(50)"`
	Email    string `valid:"MaxSize(50)"`
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

	userInfo, isExist := models.CheckAuth(username, util.EncodePwd(pwd))
	if !isExist {
		appG.Response(http.StatusOK, e.ErrorAuth, data)
		return
	}

	token, err := util.GenerateToken(userInfo.Id, username, pwd, userInfo.UserType)
	if err != nil {
		code = e.ErrorAuthToken
	} else {
		data["token"] = token
		data["id"] = userInfo.Id
		data["username"] = username
		data["email"] = userInfo.Email
		data["userType"] = userInfo.UserType
		code = e.Success
	}

	appG.Response(http.StatusOK, code, data)
}

func AdminAuth(c *gin.Context) {
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

	fmt.Println(username)
	fmt.Println(pwd)
	fmt.Println(util.EncodePwd(pwd))

	userInfo, isExist := models.CheckAuth(username, util.EncodePwd(pwd))

	if !isExist {
		appG.Response(http.StatusOK, e.ErrorAuth, data)
		return
	}

	if userInfo.UserType != models.UserTypeAdmin {
		appG.Response(http.StatusOK, e.ErrorUserNotAdmin, data)
		return
	}

	token, err := util.GenerateToken(userInfo.Id, username, pwd, userInfo.UserType)
	if err != nil {
		code = e.ErrorAuthToken
	} else {
		data["token"] = token
		data["id"] = userInfo.Id
		data["username"] = username
		data["email"] = userInfo.Email
		data["userType"] = userInfo.UserType
		code = e.Success
	}

	appG.Response(http.StatusOK, code, data)
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	email := c.PostForm("email")
	pwd := c.PostForm("pwd")

	fmt.Println(username)
	fmt.Println(email)
	fmt.Println(pwd)

	valid := validation.Validation{}
	a := auth{Username: username, Pwd: pwd, Email: email}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	isExist := models.ExistUserByUsername(username)
	if isExist {
		appG.Response(http.StatusOK, e.ErrorUsernameExists, data)
		return
	}

	cData := make(map[string]interface{})
	cData["username"] = username
	cData["pwd"] = util.EncodePwd(pwd)
	cData["email"] = email
	cData["user_type"] = models.UserTypeNormal
	userInfo, ok := models.AddUser(cData)
	if !ok {
		appG.Response(http.StatusOK, e.ErrorUserCreateFailed, data)
		return
	}

	token, err := util.GenerateToken(userInfo.Id, username, pwd, userInfo.UserType)
	if err != nil {
		code = e.ErrorAuthToken
	} else {
		data["token"] = token
		data["id"] = userInfo.Id
		data["username"] = username
		data["email"] = userInfo.Email
		data["userType"] = userInfo.UserType
		code = e.Success
	}

	appG.Response(http.StatusOK, code, data)
}
