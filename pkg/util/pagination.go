package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/kingsill/gin-example/pkg/setting"
)

// GetPage page 1  0；page 2  10；page 3  20
func GetPage(c *gin.Context) int {
	result := 0                                 //默认查询结果为0，及第1页
	page, _ := com.StrTo(c.Query("page")).Int() //查询url中包含的page信息并将其转化为int类型
	if page > 0 {                               //如果获取的页数大于0
		result = (page - 1) * setting.PageSize //返回(页数-1)×每页大小(在setting.go中已经配置过)
	}

	return result
}
