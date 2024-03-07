package v1

import (
	"github.com/astaxie/beego/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/kingsill/gin-example/models"
	"github.com/kingsill/gin-example/pkg/e"
	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/pkg/util"
)

// @Summary	查询标签
// @Produce	json
// @Param		name	query		string	true	"name"
// @Success	200		{string}	json	"{"code":200,"data":{},"msg":"ok"}"
// @Router		/api/v1/tags [get]
func GetTags(c *gin.Context) {
	//查询参数方法，及url中？name=xxx
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS //使用之前约定的错误码

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	//gin.h是一种简便的返回json的方式
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary	新增文章标签
// @Produce	json
// @Param		name		query		string	true	"Name"
// @Param		state		query		int		false	"State"
// @Param		created_by	query		int		false	"CreatedBy"
// @Success	200			{string}	json	"{"code":200,"data":{},"msg":"ok"}"
// @Router		/api/v1/tags [post]
func AddTag(c *gin.Context) {

	//参数查询 url中name
	name := c.Query("name")
	//参数查询 state 这里设置默认为0
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//参数查询
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	//运行到这里设置为 参数错误
	code := e.INVALID_PARAMS

	//将gin获得的数据与数据库作比较，进行验证
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	//json相应
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary	修改文章标签
// @Produce	json
// @Param		id			path		int		true	"ID"
// @Param		name		query		string	true	"ID"
// @Param		state		query		int		false	"State"
// @Param		modified_by	query		string	true	"ModifiedBy"
// @Success	200			{string}	json	"{"code":200,"data":{},"msg":"ok"}"
// @Router		/api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	//.param 动态参数查询，并将其确定转换为int
	id := com.StrTo(c.Param("id")).MustInt()

	//参数查询，查询对应key
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	//设定验证信息
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary	删除文章标签
// @Produce	json
// @Param		name	path		int		true	"id"
// @Success	200		{string}	json	"{"code":200,"data":{},"msg":"ok"}"
// @Router		/api/v1/tags{id} [delete]
func DeleteTag(c *gin.Context) {
	//动态参数查询
	id := com.StrTo(c.Param("id")).MustInt()

	//验证信息
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
