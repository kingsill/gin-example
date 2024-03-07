package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/kingsill/gin-example/models"
	"github.com/kingsill/gin-example/pkg/e"
	"github.com/kingsill/gin-example/pkg/logging"
	"github.com/kingsill/gin-example/pkg/util"
)

// 定义我们验证用户所需的信息，同时定义valid验证的预定信息，即一定要有并且最大字符数为50
type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary	生成用户token
// @Param		username	query		string	true	"username"
// @Param		password	query		string	true	"password"
// @Success	200			{string}	json	"{"code":200."data":{},"msg":"ok"}"
// @Router		/auth [get]
func GetAuth(c *gin.Context) {
	//参数查询，？=模式，获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	//通过设立结构体验证的valid验证，即所需并且最大为50个字符
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	//建立存储信息的map key类型为string，val类型为任意值
	data := make(map[string]interface{})

	//设置默认code为参数错误
	code := e.INVALID_PARAMS

	//通过前序验证，继续通过数据库进行验证
	if ok {
		//通过数据库进行验证
		isExist := models.CheckAuth(username, password)

		//如果通过数据库验证
		if isExist {
			//创建token令牌
			token, err := util.GenerateToken(username, password)

			if err != nil { //如果生成令牌失败
				code = e.ERROR_AUTH_TOKEN
			} else { //生成token成功，进行存储
				data["token"] = token
				code = e.SUCCESS
			}

		} else { //没通过数据库验证
			code = e.ERROR_AUTH
		}

	} else { //没通过前序验证
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	//json相应
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
