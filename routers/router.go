package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kingsill/gin-example/middleware/jwt"
	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/routers/api"
	"github.com/kingsill/gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	//注册一个新的router
	r := gin.New()

	//使用logger中间件
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	//将运行模式放到setting中设置的模式上
	gin.SetMode(setting.RunMode)

	//获取token
	r.GET("/auth", api.GetAuth)

	//路由分组，统一管理，统一增加 前缀。	在分组后通过全局路由.use注册中间件
	apiv1 := r.Group("/api/v1").Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	//将本次注册的router返回，方便使用
	return r
}
