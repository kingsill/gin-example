package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kingsill/gin-example/pkg/e"
	"github.com/kingsill/gin-example/pkg/util"
)

// 自定义中间件
func JWT() gin.HandlerFunc {
	//返回.context函数
	return func(c *gin.Context) {
		var code int
		var data interface{}

		//默认是正确状态
		code = e.SUCCESS

		//参数查询url中token关键字
		token := c.Query("token")

		//如果为空，则进行相关提示
		if token == "" {
			code = e.INVALID_PARAMS
		} else { //如果右token，进行token的解析
			claims, err := util.ParseToken(token)
			if err != nil { //如果解析出错，相关提示
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt { //如果解析出来token已过期，则也有相关提示
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		//后续处理，如果之前步骤有错误，进行一下操作
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			//放弃后续中间件的执行，即如果有错，后续中间件都不执行
			c.Abort()
			return
		}

		//如果没错，放行	next前为请求中间件，next后为相应中间件
		c.Next()
	}
}
