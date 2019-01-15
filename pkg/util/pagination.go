package util

import (
	"blog-go-server/pkg/constmap"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// 根据页码与每页数量获取sql的查询偏移量
func GetQueryOffset(pageNum int, pageLimit int) int {
	result := 0
	if pageNum > 0 {
		result = (pageNum - 1) * pageLimit
	}
	return result
}

func GetPageNum(c *gin.Context) int {
	result := 1
	pageNum, _ := com.StrTo(c.Query("page_num")).Int()
	if pageNum >= 1 {
		return pageNum
	}
	return result
}

func GetPageLimit(c *gin.Context) int {
	result := constmap.DEFAULT_PAGE_LIMIT
	pageLimit, _ := com.StrTo(c.Query("page_limit")).Int()
	if pageLimit > 0 && pageLimit <= constmap.MAX_PAGE_LIMIT {
		result = pageLimit
	}
	return result
}
