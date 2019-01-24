package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"blog-go-server/pkg/e"
	"blog-go-server/models"
	"blog-go-server/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Pwd      string `valid:"Required; MaxSize(50)"`
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	pwd := c.Query("pwd")

	valid := validation.Validation{}
	a := auth{Username: username, Pwd: pwd}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams
	if ok {
		_, isExist := models.CheckAuth(username, pwd)
		if isExist {
			token, err := util.GenerateToken(0,username, pwd)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token
				code = e.Success
			}
		} else {
			code = e.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
