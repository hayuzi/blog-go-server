package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"blog-go-server/pkg/constmap"
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
	pageNum, _ := com.StrTo(c.Query("pageNum")).Int()
	if pageNum >= 1 {
		return pageNum
	}
	return result
}

func GetPageSize(c *gin.Context) int {
	result := constmap.DefaultPageSize
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	if pageSize > 0 && pageSize <= constmap.MaxPageSize {
		result = pageSize
	}
	return result
}
