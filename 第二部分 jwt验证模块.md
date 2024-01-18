# 使用 JWT 进行身份校验
在前面几节中，我们已经基本的完成了 `API’s` 的编写，但是，还存在一些非常严重的问题，例如，我们现在的` API `是可以随意调用的，这显然还不安全全，在本文中我们通过` jwt-go （GoDoc）`的方式来简单解决这个问题。

## jwt知识点补充
[jwt官网](https://jwt.io/)
### 认识JWT 
`JSON Web Token`（**JWT**）是一个开放标准（`RFC 7519`），它定义了一种紧凑和自包含的方式，用于在各方之间作为JSON对象安全地传输信息。

作为标准，它没有提供技术实现，但是大部分的语言平台都有按照它规定的内容提供了自己的技术实现，所以实际在用的时候，只要根据自己当前项目的技术平台，到**官网上选用合适的实现库**即可。

### TOKEN是什么
`Token`，其实就是服务端生成的一串**加密字符串**、以作客户端进行请求的一个“令牌”

![Alt text](token.png)

### jwt的使用场景
以下是JWT两种使用场景：

>**授权**：这是使用` JWT` 的最常见的使用场景。用户登录后，每个后续请求都将包含 `JWT`，允许用户访问使用该令牌允许的路由、服务和资源。**单点登录**是当今广泛使用 `JWT` 的一项功能，因为它的开销很小，并且能够跨不同域轻松使用。

>**信息交换**：`JWT`是在各方之间安全传输信息的比较便捷的方式。由于 `JWT` 可以签名（例如，使用`公钥`/`私钥`对），因此可以确定发送者是否是在您的授权范围之内。并且，由于签名是使用标头和有效负载计算的，因此还可以**验证内容是否未被篡改**。

### jwt的组成
这是一个JWT的token串：
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
其实这一串是经过加密之后的密文字符串，中间通过`.`来分割。每个`.`之前的字符串分别表示`JWT`的三个组成部分：`Header`、`Payload`、`Signature`。
#### header
Header的主要作用是用来标识,通常是**两部分组成**：
>`typ`：`type` 的简写，令牌类型，也就是`JWT`。

>`alg`：`Algorithm` 的简写，加密签名算法。一般使用`HS256`，`jwt`官网提供了12种的加密算法

然后通过base64编码，将明文编码,防止在传输过程中能直接一眼看出明文并符合多种传输协议

#### payload
也称为`JWT claims`

`payload`用来承载要传递的数据，它的`json`结构实际上是对`JWT`要传递的数据的一组声明，这些声明被`JWT`标准称为`claims`，它的一个“**属性值对**”其实就是一个`claim`，每一个`claim`的都代表特定的含义和作用

claims有三类：

- **保留claims**：主要包括`iss`发行者、`exp`过期时间、`sub`主题、`aud`用户等。

key|name|说明
-|-|-
iss|发送者|	标识颁发 `JWT` 的发送主体
sub|	主题|	标识 `JWT` 的主题
aud|	接收者|标识 `JWT` 所针对的接收者。每个在处理 `JWT `的主体都必须使用受众声明中的值来标识自己。如果处理的主体在存在此声明时未将自己标识为声明中的值，则必须拒绝` JWT`
exp|到期时间|	标识不得接受 `JWT` 进行处理的过期时间。该值必须是日期类型，而且是`1970-01-01 00：00：00Z` 之后的日期秒。
nbf|	`jwt`的开始处理的时间|	标识 `JWT `开始接受处理的时间。该值必须是日期。
iat|	`jwt`发出的时间|标识 `JWT` 的发出的时间。该值必须是日期。
jti|	jwt id|令牌的区分大小写的唯一标识符，即使在不同的颁发者之间也是如此。

保留`claim`为`jwt`标准中规定的`claim`，验证方式已经定义好
- **公共claims**：定义新创的信息，比如用户信息和其他重要信息。(使用较少)
- **私有claims**：用于发布者和消费者都同意以私有的方式使用的信息。

明文实例：
```
{
  "sub": "12344321",
  "name": "Mars酱", // 私有claims
  "iat": 1516239022
}
```

base64加密后：
```
eyJzdWIiOiIxMjM0NDMyMSIsIm5hbWUiOiJNYXJz6YWxIiwiaWF0IjoxNTE2MjM5MDIyfQ
```


#### signature
`Signature `部分是对`Header`和`Payload`两部分的签名，作用是**防止 `JWT `被篡改**。这个部分的生成规则主要是是公式（伪代码）是：
```
Header中定义的签名算法alg(
    base64编码(header) + "." + base64编码(payload),
    secret//在服务端加密使用的密钥
)
```

`JWT`如果从字面上理解感觉是基于`JSON`格式用于**网络传输**的**令牌**。实际上，`JWT`是一种紧凑的`Claims`声明格式，，常见的场景如HTTP授权请求头参数和URI查询参数。`JWT`会把`Claims`转换成`JSON`格式，而这个`JSON`内容将会应用为`JWS`结构的有效载荷或者应用为`JWE`结构的（加密处理后的）原始字符串，通过消息认证码（`Message Authentication Code`或者简称`MAC`）和/或者加密操作对`Claims`进行数字签名或者完整性保护。


## 下载依赖包
```shell
go get -u github.com/dgrijalva/jwt-go
```
## 编写 jwt 工具包
我们需要编写一个`jwt`的工具包，我们在`pkg`下的`util`目录新建`jwt.go`，写入文件内容:
```go
package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/kingsill/gin-example/pkg/setting"
)

// 加载配置文件中设置的密钥
var jwtSecret = []byte(setting.JwtSecret)

// Claims 定义claims结构体
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	//创建 CustomClaims 结构体，用来封装 jwt 信息
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//创建 header和payload部分
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//得到完整的token字符串，这里为加入签名signature部分
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//解码过程
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	//验证是否时间过期
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

```
在这个工具包，我们涉及到:
- `NewWithClaims(method SigningMethod, claims Claims)`，`method`对应着`SigningMethodHMAC struct{}`，其包含`SigningMethodHS256、SigningMethodHS384、SigningMethodHS512`三种`crypto.Hash`方案
- `func (t *Token) SignedString(key interface{})` 该方法内部生成签名字符串，再用于获取完整、已签名的`token`
- `func (p *Parser) ParseWithClaims` 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回`*Token`
- `func (m MapClaims) Valid()` 验证基于时间的声明`exp, iat, nbf`，注意如果没有任何声明在令牌中，仍然会被认为是有效的。并且对于时区偏差没有计算方法

## jwt中间件编写
[中间件相关知识](https://blog.csdn.net/kingsill/article/details/133808879#t11)
[自定义中间件](https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/)
有了jwt工具包，接下来我们要编写要用于Gin的中间件，我们在middleware下新建jwt目录，新建jwt.go文件，写入内容：
```go
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
```

## 如何获取`token`  编写获取`token`的Api
那么我们如何调用它呢，我们还要获取`Token`呢？
### models逻辑编写
在`models`下新建`auth.go`文件，写入内容：
```go
package models

// jwt验证的 数据库相关操作

// Auth 用户对应的struct模型
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 根据用户名和密码查询用户是否存在
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
```

}
### 路由逻辑编写
在`routers`下的`api`目录新建`auth.go`文件，写入内容：
```go
package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/kingsill/gin-example/models"
	"github.com/kingsill/gin-example/pkg/e"
	"github.com/kingsill/gin-example/pkg/util"
)

// 定义我们验证用户所需的信息，同时定义valid验证的预定信息，即一定要有并且最大字符数为50
type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

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
			log.Println(err.Key, err.Message)
		}
	}

	//json相应
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
```

### 修改路由逻辑
我们打开`routers`目录下的`router.go`文件，修改文件内容（新增获取 `token` 的方法）：
增添一句`	r.GET("/auth", api.GetAuth)`，放在路由组之外
```go
func InitRouter() *gin.Engine {
r := gin.New()

r.Use(gin.Logger())

r.Use(gin.Recovery())

gin.SetMode(setting.RunMode)

r.GET("/auth", api.GetAuth)//------------------------

apiv1 := r.Group("/api/v1")
{
...
}

return r
}
```

## 验证token
获取`token`的` API` 方法就到这里啦，让我们来测试下是否可以正常使用吧！

重启服务后，用`GET`方式访问`http://127.0.0.1:8000/auth?username=test&password=test123456` ，查看返回值是否正确
```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE3MDUzMzk2MzUsImlzcyI6Imdpbi1ibG9nIn0.4zblfic9MdOvrg4TF9Li8nfw3FSBq3rGgKqnJnDFXYY"
    },
    "msg": "ok"
}
```
我们有了`token`的 `API`，也调用成功了

### 将中间件接入`Gin`
修改路由分组模块语句，`apiv1 := r.Group("/api/v1").Use(jwt.JWT())`，将jwt验证加入全局路由
```go
...
    apiv1 := r.Group("/api/v1")
apiv1.Use(jwt.JWT())
{
...
}
...
```
当前目录结构
```
go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
│   └── jwt
│       └── jwt.go
├── models
│   ├── article.go
│   ├── auth.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── routers
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```
到这里，我们的JWT编写就完成啦！

### 功能验证模块

我们来测试一下，再次访问

- http://127.0.0.1:8000/api/v1/articles
- http://127.0.0.1:8000/api/v1/articles?token=23131
正确的反馈应该是
```
{
  "code": 400,
  "data": null,
  "msg": "请求参数错误"
}

{
  "code": 20001,
  "data": null,
  "msg": "Token鉴权失败"
}

```

我们需要访问`http://127.0.0.1:8000/auth?username=test&password=test123456` ，得到token


再用包含`token`的 `URL` 参数去访问我们的应用 `API`，

- 这里的问题即为创建的token还需要自己复制粘贴，不能自动取用等

访问`http://127.0.0.1:8000/api/v1/articles?token=eyJhbGci...` ，检查接口返回值

```
{
    "code": 200,
    "data": {
        "lists": [],
        "total": 0
    },
    "msg": "ok"
}
```
验证正确，文章列表取决于数据库内容