package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"test2/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page < 0 || page == 0 {
		page = 1
	}
	if page > 0 {
		result = (page-1)*setting.PageSize
	}
	return result
}
