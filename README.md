# go-gin-example 
从2023.12.22开始学习煎鱼大佬的该项目
[煎鱼大佬网站](https://eddycjy.com/go-categories/)
[原github地址](https://github.com/eddycjy/go-gin-example)

# 环境准备
## 初始化 Go Modules
**Go Modules** 是go的**依赖包管理**工具,现在的go版本自动打开modules，目前的`go get`命令也是需要进行初始化才能进行拉取

在准备的文件夹的终端中执行
```
$ go env -w GOPROXY=https://goproxy.cn,direct

$ go mod init [github.com/kingsill/gin-example]
```
`go env -w GOPROXY=...`：设置 GOPROXY 代理，这里主要涉及到两个值，第一个是 https://goproxy.cn，它是由七牛云背书的一个强大稳定的 Go 模块代理，可以有效地解决你的外网问题；第二个是 direct，它是一个特殊的 fallback 选项，它的作用是用于指示 Go 在拉取模块时遇到错误会回源到模块版本的源地址去抓取（比如 GitHub 等）。

`go mod init [MODULE_PATH]`：初始化 Go modules，它将会生成 go.mod 文件，需要注意的是 MODULE_PATH 填写的是模块引入路径，你可以根据自己的情况修改路径。
这里我们使用github域名作为项目名是证明这个包使用github进行存储。


此时 **.mod**文件内容。是当前的 **模块路径** 和 预期的 **Go 语言版本** 。
```
module github.com/kingsill/gin-example

go 1.21.4
```



### 基础使用
- 用 go get 拉取新的依赖
    - 拉取最新的版本(优先择取 tag)：go get golang.org/x/text@latest
    - 拉取 master 分支的最新 commit：go get golang.org/x/text@master
    - 拉取 tag 为 v0.3.2 的 commit：go get golang.org/x/text@v0.3.2
    - 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：go get golang.org/x/text@342b2e
    - 用 go get -u 更新现有的依赖
    - 用 go mod download 下载 go.mod 文件中指明的所有依赖
    - 用 go mod tidy 整理现有的依赖
    - 用 go mod graph 查看现有的依赖结构
    - 用 go mod init 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)
- 用 go mod edit 编辑 go.mod 文件
- 用 go mod vendor 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)
- 用 go mod verify 校验一个模块是否被篡改过

## gin 安装
在项目的根目录的命令行执行
```
$ go get -u github.com/gin-gonic/gin
```

此时我们的目录如下所示，**多出.sum文件**
`go.sum `文件详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。
```
gin-example learn
├─ .idea
│  └─ workspace.xml
├─ go.mod
├─ go.sum
└─ README.md

```

此时go.mod文件也多了一些内容。

```
module github.com/kingsill/gin-example

go 1.21.4

require (
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

**go.mod** 文件是启用了 Go modules 的项目所必须的最重要的文件，因为它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

你可能还会疑惑 indirect 是什么东西，indirect 的意思是传递依赖，也就是非直接依赖。

### 测试gin是否引入
[gin部分学习博客](https://blog.csdn.net/kingsill/article/details/133611318?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522170324611116800192214773%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fblog.%2522%257D&request_id=170324611116800192214773&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~blog~first_rank_ecpm_v1~rank_v31_ecpm-1-133611318-null-null.nonecase&utm_term=gin&spm=1018.2226.3001.4450)

编写test.go文件并执行
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	e.Run()
}

```
访问 `127.0.0.1:8080`,如下则正确安装
![Alt text](test-gin.png)

此时我们go.mod文件中使用的依赖 后缀也 都是indirect，这时候我们需要使用`go mod tidy`进行依赖整理，这个命令非常常用

整理完之后go.mod文件如下所示：gin已经是直接依赖
```go
module go-gin-example

go 1.21.4

require github.com/gin-gonic/gin v1.9.1

require (
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

# gin搭建Blog API's
由于要避免以下问题：
- 程序的文本配置写在代码中
- API的错误码硬编码在程序中
- DB句柄谁都open，没有通一管理
- 获取分页等公共参数 一人一套逻辑
我们选择进行配置文件，这里我们选择[go-ini/ini](https://github.com/go-ini/ini),要先简单阅读其[中文文档](https://ini.unknwon.io/)

 ## go-ini
 ### 简述配置文件
配置文件本质上是包含成功操作程序所需信息的文件，这些信息以特定方式构成。是用户可配置的，通常存储在纯文本文件中

配置文件可以是各种格式，完全凭借程序员的发挥，不过出于方便，大部分会选择的配置文件格式集中在那几种.一般而言程序启动时，会加载该程序对应的配置文件内的信息

**有这么几点作用**：
1.数据库的连接工作
2.端口号的配置
3.打印日志等等

## 阶段目标 编写简单API错误码包 完成一个demo

### 初始化项目
将目录结构更新到如下所示：

```
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据

### 初始化项目数据库
新建 `blog` 数据库，编码为`utf8_general_ci`，在 blog 数据库下，新建以下表
1. 标签表
```sql
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```

2. 文章表

```sql
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
```

3. 认证表

```sql
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```

### 编写项目配置包

#### 拉取go-ini配置包
	```
 	go get -u github.com/go-ini/ini
	```

#### 在conf目录下新建app.ini文件，写入内容：
	[ini中文文档](https://ini.unknwon.io/docs/intro/getting_started)

	```ini
	#debug or release
	RUN_MODE = debug

	[app]
	PAGE_SIZE = 10
	JWT_SECRET = 23347$040412

	[server]
	HTTP_PORT = 8000
	READ_TIMEOUT = 60
	WRITE_TIMEOUT = 60

	[database]
	TYPE = mysql
	USER = 数据库账号
	PASSWORD = 数据库密码
	#127.0.0.1:3306
	HOST = 数据库IP:数据库端口号
	NAME = blog
	TABLE_PREFIX = blog_
	```
	**相关内容记得填写自己的配置进去**
	可以看到`ini配置`文件中，`默认分区`中定义RUN_MODE为debug，`app分区`中定义两项，`server分区`定义HTTP_PORT等，`database分区`定义了数据库相关信息
	>个人理解：ini配置文件将我们可能用到的配置信息同意展示在配置文件里，有助于后续管理及更新

#### 建立调用配置的`setting` 模块
	go-gin-example的`pkg`目录下新建`setting目录`，新建 `setting.go` 文件，写入内容：

	```go
	package setting

	import (
	"log"
	"time"

	"github.com/go-ini/ini"
	)

	var (
	Cfg *ini.File //加载配置文件读取

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
	)

	func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini") //加载cong/app.ini文件，即我们自己创建的配置文件
	if err != nil {                     //错误提示
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	}

	// LoadBase 加载运行模式
	func LoadBase() {
	//默认分区，使用""，key值为RUN_MODE，若配置中未指定，则默认为debug
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	}

	// LoadServer 加载服务器相关配置
	func LoadServer() {
		sec, err := Cfg.GetSection("server") //从配置文件检索server分区的内容，sec为获取的指定分区
		if err != nil {                      //错误提示
			log.Fatalf("Fail to get section 'server': %v", err)
		}

		//设置服务器配置，如果配置中未指定则使用默认值。
		HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
		ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
		WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	}

	// LoadApp 从配置文件中加载应用程序特定的设置
	func LoadApp() {
		sec, err := Cfg.GetSection("app") //从配置文件检索app分区的内容，sec为获取的指定分区
		if err != nil {                   //错误提示
			log.Fatalf("Fail to get section 'app': %v", err)
		}
		
		// 设置应用程序配置，如果配置中未指定则使用默认值。
		JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
		PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	}

	```

	当前目录结构
	```
	go-gin-example
	├── conf
	│   └── app.ini
	├─middleware
	├─models
	├─pkg
	│  └─setting
	|       └──setting.go
	├─routers
	└─runtime
	```

#### 编写 API 错误码包
建立错误码的`e`模块，在`go-gin-example`的`pkg`目录下新建`e目录`，新建`code.go`和`msg.go`文件，写入内容：

>这里即创建golang中的枚举方法，便于后期处理和维护，[GO GORM 自定义数据类型-枚举](https://blog.csdn.net/kingsill/article/details/134867309?spm=1001.2014.3001.5502)这边文章具体提及其方法，可以参考

1. code.go
	在code.go中 定义 具有意义的**字符** 为 对应的**错误码**，具有编码作用
	```go
	package e

	const (
		SUCCESS = 200
		ERROR = 500
		INVALID_PARAMS = 400

		ERROR_EXIST_TAG = 10001
		ERROR_NOT_EXIST_TAG = 10002
		ERROR_NOT_EXIST_ARTICLE = 10003

		ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
		ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
		ERROR_AUTH_TOKEN = 20003
		ERROR_AUTH = 20004
	)
	```

2. msg.go
	在`msg.go`中 通过建立 int对应string的**map表** 将 code.go 中的**字符常量** 对应为 **字符串**，之后 通过`GetMsg`函数 可以直接将 **错误码** **对应**到 想要传递的**错误信息**

	```go
	package e

	var MsgFlags = map[int]string {
		SUCCESS : "ok",
		ERROR : "fail",
		INVALID_PARAMS : "请求参数错误",
		ERROR_EXIST_TAG : "已存在该标签名称",
		ERROR_NOT_EXIST_TAG : "该标签不存在",
		ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
		ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
		ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
		ERROR_AUTH_TOKEN : "Token生成失败",
		ERROR_AUTH : "Token错误",
	}

	func GetMsg(code int) string {
		msg, ok := MsgFlags[code]
		if ok {
			return msg
		}

		return MsgFlags[ERROR]
	}
	```


#### 编写工具包
在`go-gin-example`的`pkg`目录下新建`util`目录，并拉取`com`的依赖包，如下：
```
$ go get -u github.com/unknwon/com
```

##### 编写分页页码的获取方法
在`util`目录下新建`pagination.go`，写入内容：
```go
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

```


关于其中gin的查询参数部分可以参考这部分[gin 查询参数](https://blog.csdn.net/kingsill/article/details/133611318?spm=1001.2014.3001.5502#t18)


#### 编写 models init
//配置部分数据库连接相关内容
拉取`gorm`的依赖包，拉取`mysql驱动`的依赖包，如下：
```
$ go get -u github.com/jinzhu/gorm
$ go get -u github.com/go-sql-driver/mysql
```

有关gorm和mysql可以参看这两部分文章：
[gorm](https://blog.csdn.net/kingsill/category_12535358.html?spm=1001.2014.3001.5482)			[mysql](https://blog.csdn.net/kingsill/category_12535359.html) 

完成后，在`go-gin-example`的`models`目录下新建`models.go`，用于models的初始化使用

```go
package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kingsill/gin-example/pkg/setting"
)

// 定义一个全局的数据库连接变量
var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 将一下定义为init函数
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	//加载配置文件中database分区的数据
	sec, err := setting.Cfg.GetSection("database") //cfj在setting模块中已经通过init函数进行初始化
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//配置导入
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//使用gorm框架初始化数据库连接
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	//自定义默认表的表名，使用匿名函数，在原默认表名的前面加上配置文件中定义的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//gorm默认使用复数映射，当前设置后即进行严格匹配
	db.SingularTable(true)
	//log记录打开
	db.LogMode(true)

	//进行连接池设置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB 与数据库断开连接函数
func CloseDB() {
	defer db.Close()
}

```
此时我依旧不太明白中间闭包的作用，还希望路过大神能够解惑

### 编写项目启动、路由文件 Demo

在go-gin-example下建立main.go作为启动文件（也就是main包），我们先写个**Demo**，帮助大家理解，写入文件内容：
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kingsill/gin-example/pkg/setting"
)

// 这里的测试demo只使用到了pkg/setting包，即读取配置文件app. ini文件部分
func main() {

	//创建一个默认路由器router
	router := gin.Default()

	//创建一个对应的路由handler处理get请求，这里定义了测试的ip即x.x.x/test
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	//创建一个http服务器，将前面的router绑定为这里的处理器
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//启动服务器
	s.ListenAndServe()
}
```

执行go run main.go，查看命令行是否显示:
```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> main.main.func1 (3 handlers)
```

在本机执行curl 127.0.0.1:8000/test,或者在浏览器中输入ip地址，检查是否返回{"message":"test"}。

笔者的路由、handler等定义可能比较混乱，也希望大家指出错误
### 知识点部分
那么，我们来延伸一下 Demo 所涉及的知识点！

**标准库**
- fmt：实现了类似 C 语言 printf 和 scanf 的格式化 I/O。格式化动作（‘verb’）源自 C 语言但更简单
net/http：提供了 HTTP 客户端和服务端的实现

**Gin**
- gin.Default()：返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
- router.GET(…){…}：创建不同的 HTTP 方法绑定到Handlers中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
- gin.H{…}：就是一个map[string]interface{}
- gin.Context：Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等

**&http.Server 和 ListenAndServe？**

1. http.Server：
```go
type Server struct {
    Addr    string
    Handler Handler
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```
- **Addr**：监听的 TCP 地址，格式为:8000
- **Handler**：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
- **TLSConfig**：安全传输层协议（TLS）的配置
- **ReadTimeout**：允许读取的最大时间
- **ReadHeaderTimeout**：允许读取请求头的最大时间
- **WriteTimeout**：允许写入的最大时间
- **IdleTimeout**：等待的最大时间
- **MaxHeaderBytes**：请求头的最大字节数
- **ConnState**：指定一个可选的回调函数，当客户端连接发生变化时调用
- **ErrorLog**：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）

2. ListenAndServe：
```go
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```
开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求。

我们在源码中看到Addr是调用我们在&http.Server中设置的参数，因此我们在设置时要用&，我们要改变参数的值，因为我们ListenAndServe和其他一些方法需要用到&http.Server中的参数，他们是相互影响的。

3. http.ListenAndServe和 连载一 的r.Run()有区别吗？
我们看看`r.Run`的实现：
```go
func (engine *Engine) Run(addr ...string) (err error) {
    defer func() { debugPrintError(err) }()

    address := resolveAddress(addr)
    debugPrint("Listening and serving HTTP on %s\n", address)
    err = http.ListenAndServe(address, engine)
    return
}
```
通过分析源码，得知**本质上没有区别**，同时也得知了启动gin时的监听 debug 信息在这里输出。

### 最后的小改进部分
 Demo 的router.GET等路由规则可以不写在main包中吗？
 我们发现router.GET等路由规则，在 Demo 中被编写在了main包中，感觉很奇怪，我们去抽离这部分逻辑！

在go-gin-example下routers目录新建router.go文件，写入内容：
```go
package routers

import (
    "github.com/gin-gonic/gin"

    "github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {

	//创建新的router而不是默认。
	r := gin.New()
	
	//定义router的配置
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.RunMode)

	//将处理函数搬到这里
    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })

    return r
}
```

**修改main.go文件内容**
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/routers"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
```

**当前目录结构：**
```
go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   └── models.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   └── router.go
├── runtime
```


## 编写tag的api
目标：完成博客的标签类接口定义和编写

### 定义接口
本节正是编写标签的逻辑，我们想一想，一般接口为增删改查是基础的，那么我们定义一下接口吧！
- 获取标签列表：GET("/tags”)
- 新建标签：POST("/tags”)
- 更新指定标签：PUT("/tags/:id”)
- 删除指定标签：DELETE("/tags/:id”)

### 编写路由空壳
开始编写路由文件逻辑，在routers下新建api目录，我们当前是第一个 API 大版本，因此在api下新建v1目录，再新建tag.go文件，写入内容：
```go
package v1

import (
    "github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```

### 注册路由
我们打开routers下的router.go文件，修改文件内容为：
```go
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


```

### 检查路由是否注册成功
回到命令行，执行`go run main.go`，检查路由规则是否注册成功。

```go
$ go run main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/tags              --> gin-blog/routers/api/v1.GetTags (3 handlers)
[GIN-debug] POST   /api/v1/tags              --> gin-blog/routers/api/v1.AddTag (3 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> gin-blog/routers/api/v1.EditTag (3 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> gin-blog/routers/api/v1.DeleteTag (3 handlers)
```

### 下载依赖包
首先我们要拉取`validation`的依赖包，在后面的接口里会使用到表单验证
```
$ go get -u github.com/astaxie/beego/validation
```

### 编写标签列表的 models 逻辑
创建`models`目录下的`tag.go`，写入文件内容：
```go
package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Tag 定义tag表的相关表头
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//以下两个函数都为命名返回，函数内部有相关定义

// GetTags page-size,每页显示的tag数	pageNum，即为从第几条记录开始显示，从0开始计数
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {

	// where查询使用map条件
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags) //db在models.go中已经定义，同一个包可以直接调用
	return
}

// GetTagTotal 查询tag总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}
```
1. 我们创建了一个Tag struct{}，用于Gorm的使用。并给予了附属属性json，这样子在c.JSON的时候就会自动转换格式，非常的便利

2. 可能会有的初学者看到return，而后面没有跟着变量，会不理解；其实你可以看到在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用，因为他在一开始就被声明了

3. 有人会疑惑db是哪里来的；因为在同个models包下，因此db *gorm.DB是可以直接使用的

### 编写标签列表的路由逻辑
打开`routers`目录下 `v1` 版本的`tag.go`，第一我们先编写获取标签列表的接口

修改文件内容：

```go
package v1

import (
    "net/http"

    "github.com/gin-gonic/gin"
    //"github.com/astaxie/beego/validation"
    "github.com/Unknwon/com"

    "gin-blog/pkg/e"
    "gin-blog/models"
    "gin-blog/pkg/util"
    "gin-blog/pkg/setting"
)

// GetTags 获取多个文章标签
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

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```

1. c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值
2. code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
3. util.GetPage保证了各接口的page处理是一致的
4. c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应

在本机执行`curl 127.0.0.1:8000/api/v1/tags`，正确的返回值为`{"code":200,"data":{"lists":[],"total":0},"msg":"ok"}`，若存在问题请结合` gin `结果进行拍错。

在获取标签列表接口中，我们可以根据`name、state、page`来筛选查询条件，分页的步长可通过`app.ini`进行配置，以`lists、total`的组合返回达到分页效果。

### 编写新增标签的 models 逻辑
接下来我们编写新增标签的接口

打开`models`目录下的`tag.go`，修改文件（增加 2 个方法）：
```go
...
// ExistTagByName 根据name查询tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	//查询词条
	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// AddTag 创建新tag
func AddTag(name string, state int, createdBy string) bool {
	//在对应数据表中创建词条
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

...
```

### 编写新增标签的路由逻辑
打开`routers`目录下的`tag.go`，修改文件（变动 `AddTag` 方法）：

```go
package v1

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    "github.com/Unknwon/com"

    "gin-blog/pkg/e"
    "gin-blog/models"
    "gin-blog/pkg/util"
    "gin-blog/pkg/setting"
)

...

// AddTag 新增文章标签
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
...
```
用`Postman`用 `POST `访问`http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test`，查看`code`是否返回`200`及`blog_tag`表中是否有值，有值则正确。

### 编写 models callbacks
但是这个时候大家会发现，我明明新增了标签，但`created_on`居然没有值，那做修改标签的时候`modified_on`会不会也存在这个问题？

为了解决这个问题，我们需要打开`models`目录下的`tag.go`文件，修改文件内容（修改包引用和增加 2 个方法）：

```go
package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

...

// BeforeCreate 建立hook钩子函数，在创建之前插入时间值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	//我们在定义的时候createon是int类型，因此我们这里使用unix方法将时间转变为时间戳
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate 同为hook钩子函数，更新的时候加入修改时间值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
...
```

重启服务，再在用`Postman`用 `POST` 访问`http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test`，发现`created_on`已经有值了！

**这里涉及到gorm的相关方式，更多的可以查看这篇[gormhook相关的文章](https://blog.csdn.net/kingsill/article/details/134607359?spm=1001.2014.3001.5501)**

### 编写其余接口的路由逻辑
接下来，我们一口气把剩余的两个接口（`EditTag、DeleteTag`）完成吧

打开`routers`目录下 `v1` 版本的`tag.go`文件，修改内容：
```go
...
// EditTag 修改文章标签
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

// DeleteTag 删除文章标签
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

```

### 编写其余接口的 models 逻辑
打开`models`下的`tag.go`，修改文件内容：


```go
...

// ExistTagByID 根据id查询表中tag是否存在
func ExistTagByID(id int) bool {
	var tag Tag //实例化tag

	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// DeleteTag 根据id删除表中tag
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

// EditTag 修改tag
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

...

```

### 验证功能
重启服务，用` Postman`

- `PUT` 访问 `http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1` ，查看 `code` 是否返回 `200`
- `DELETE` 访问 `http://127.0.0.1:8000/api/v1/tags/1` ，查看 `code` 是否返回 `200`
至此，`Tag` 的 `API’s` 完成，下一节我们将开始 `Article` 的 `API’s` 编写！

## 完成博客的文章类接口定义和编写
### 定义接口
本节编写文章的逻辑，我们定义一下接口吧！

- **获取文章列表**：GET("/articles”)
- **获取指定文章**：POST("/articles/:id”)
- **新建文章**：POST("/articles”)
- **更新指定文章**：PUT("/articles/:id”)
- **删除指定文章**：DELETE("/articles/:id”)

### 编写路由空壳
在`routers`的 `v1` 版本下，新建`article.go`文件，写入内容：
```go
package v1

import (
    "github.com/gin-gonic/gin"
)

//获取单个文章
func GetArticle(c *gin.Context) {
}

//获取多个文章
func GetArticles(c *gin.Context) {
}

//新增文章
func AddArticle(c *gin.Context) {
}

//修改文章
func EditArticle(c *gin.Context) {
}

//删除文章
func DeleteArticle(c *gin.Context) {
}

```
我们打开`routers`下的`router.go`文件，修改文件内容为：
```go
package routers

import (
    "github.com/gin-gonic/gin"
	
	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
    ...
    apiv1 := r.Group("/api/v1")
    {
        ...
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

    return r
}

```
当前目录结构：
```
go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime

```
在基础的路由规则配置结束后，我们开始编写我们的接口吧！

### 编写models逻辑
创建`models`目录下的`article.go`，写入文件内容：
```go
package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

// Article 建立对应article表的struct结构体，方便进行信息读写
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate 与tag的逻辑相同 为了插入createOn时间戳
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate 与tag的逻辑相同 是为了更新数据是插入modifiedOn数据
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

```
我们创建了一个`Article struct {}`，与`Tag`不同的是，`Article`多了几项，如下：
- `gorm:index`，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
- `Tag`字段，实际是一个嵌套的`struct`，它利用`TagID`与`Tag`模型相互关联，在执行查询的时候，能够达到`Article、Tag`关联查询的功能
- `time.Now().Unix()` 返回当前的时间戳
接下来，请确保已对上一章节的内容通读且了解，由于逻辑偏差不会太远，我们本节直接编写这五个接口


