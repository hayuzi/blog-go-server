package jwt

import (
	"blog-go-server/models"
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		var code int
		data := make(map[string]interface{})

		code = e.Success
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			appG.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}
		c.Next()
	}
}

func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		var code int
		data := make(map[string]interface{})

		code = e.Success
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			} else if claims.UserType != models.UserTypeAdmin {
				code = e.ErrorAuthCheckTokenFail
			}
		}

		if code != e.Success {
			appG.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}
		c.Next()
	}
}
