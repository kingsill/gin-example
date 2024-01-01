package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	//注册一个新的router
	r := gin.New()

	//使用logger
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//将运行模式放到setting中设置的模式上
	gin.SetMode(setting.RunMode)

	//路由分组，统一管理，统一增加 前缀
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	//将本次注册的router返回，方便使用
	return r
}
