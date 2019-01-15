package util

import (
	"blog-go-server/pkg/constmap"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// 根据页码与每页数量获取sql的查询偏移量
func GetQueryOffset(pageNum int, pageSize int) int {
	result := 0
	if pageNum > 0 {
		result = (pageNum - 1) * pageSize
	}
	return result
}

func GetPageNum(c *gin.Context) int {
	result := constmap.DefaultPageNum
	pageNum, _ := com.StrTo(c.Query("page_num")).Int()
	if pageNum >= 1 {
		return pageNum
	}
	return result
}

func GetPageSize(c *gin.Context) int {
	result := constmap.DefaultPageLimit
	pageSize, _ := com.StrTo(c.Query("page_limit")).Int()
	if pageSize > 0 && pageSize <= constmap.MaxPageLimit {
		result = pageSize
	}
	return result
}
