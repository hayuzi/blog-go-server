package common

import (
	"blog-go-server/pkg/util"
	"errors"
	"time"
)

// 获取jwt登陆用户的基础信息
func getLoginClaims(token string) (*util.Claims, error) {
	var claims = new(util.Claims)
	var err error
	if token != "" {
		claims, err = util.ParseToken(token)
		if err == nil && time.Now().Unix() > claims.ExpiresAt {
			err = errors.New("token expired")
		}
	}
	return claims, nil
}
