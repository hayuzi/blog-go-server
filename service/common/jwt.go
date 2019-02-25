package common

import (
	"blog-go-server/pkg/util"
	"github.com/gin-gonic/gin"
)

// 获取jwt登陆用户的基础信息
func GetLoginClaims(c *gin.Context) (*util.Claims, error) {
	token := c.Query("token")
	var claims = new(util.Claims)

	//if token == "" {
	//	return claims, errors.New("token empty")
	//}

	claims, _ = util.ParseToken(token)
	//if err != nil {
	//	return claims, errors.New(e.GetMsg(e.ErrorAuthCheckTokenFail))
	//} else if time.Now().Unix() > claims.ExpiresAt {
	//	return claims, errors.New(e.GetMsg(e.ErrorAuthCheckTokenTimeout))
	//}

	return claims, nil
}
