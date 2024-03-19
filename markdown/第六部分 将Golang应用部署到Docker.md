# Docker
## docker相关部分知识
[runoob docker教程](https://www.runoob.com/docker/docker-tutorial.html)
[同站友人关于docker的相关介绍](https://blog.csdn.net/weixin_42592282/article/details/121783396)

简而言之，`docker`是一款轻量级的虚拟机

`Docker` 是一个用于开发，交付和运行应用程序的开放平台。`Docker` 使您能够将应用程序与基础架构分开，从而可以快速交付软件。借助` Docker`，您可以与管理应用程序相同的方式来管理基础架构。通过利用 Docker 的方法来快速交付，测试和部署代码，您可以大大减少编写代码和在生产环境中运行代码之间的延迟

## docker安装
### winddows环境
>由于`Docker`并非通用的容器工具，其**依赖于已经存在并运行的`linux`内核环境**

`Docker` 实质上是在已经运行的 `Linux` 下制造了一个隔离的文件环境，因此它执行的效率几乎等同于所部署的 Linux 主机。

因此，`Docker` 必须部署在 Linux 内核的系统上。如果其他系统想部署 `Docker` 就**必须安装一个虚拟 Linux 环境**。

目前博主推荐在`windows`使用`wsl2`作为linux环境
[这里再次引用同站博主的wsl安装教程](https://blog.csdn.net/wojiuguowei/article/details/122100090?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522171074604616800222884664%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=171074604616800222884664&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-1-122100090-null-null.142^v99^pc_search_result_base2&utm_term=wsl&spm=1018.2226.3001.4187)

之后从[docker官方](https://docs.docker.com/desktop/install/windows-install/)下载安装`docker Desktop`即可

更多docker相关知识就请大家自己学习了

# 本部分目标
将`go-gin-example`部署到`docker`中

# 实现
## 编写dockerfile
在我们的项目根目录下创建dockerfile文件，写入以下内容
```dockerfile
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/kingsill/gin-example
COPY . $GOPATH/src/github.com/kingsill/gin-example

#RUN go build main.go
RUN go build .

EXPOSE 8000

#ENTRYPOINT ["./main"]
ENTRYPOINT ["./gin-example"]
```
- `FROM golang:latest`

指定使用最新版的官方的`Golang`镜像作为基础镜像，用于构建`docker容器`

- `ENV GOPROXY https://goproxy.cn,direct`

对创建容器的**全局变量**进行设置，设置go模块代理环境为**国内镜像**

- `WORKDIR $GOPATH/src/github.com/kingsill/gin-example`

设置容器内的**工作目录**，如果不存在会自动创建，首次出现，为进入容器后默认的工作目录，不需要cd

- `COPY . $GOPATH/src/github.com/kingsill/gin-example`

将当前目录的所有内容（.指代所有文件）**复制**到指定目录（刚才设置的容器内的工作目录）

- `RUN go build .`

Run命令等同于在终端操作的shell命令，**只在build时执行**。这里即相当于在创建的容器的工作目录中编译当前文件夹内的go文件，**生成的exe文件名字为根目录的名称**

- `EXPOSE 8000`

expose声明对外暴露的端口，可以不写，`docker run -p`可以指定暴露端口

- `ENTRYPOINT ["./gin-example"]`

e`ntrypoint`为容器**启动时执行的命令**，这里即运行build生成的可执行文件，让程序随容器运行而启动

## 拉取mysql镜像
1. 拉取镜像
从Docker公共仓库下载MySql镜像（国内建议先配置镜像源）
```shell
docker pull mysql
```
2. 创建、运行MySql容器
```shell
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -d mysql
```
> `--name`：指定容器名称

> `-p 3306：3306`：指定端口映射，将容器内部的Mysql服务端口3306映射到主机的3306端口。

> `-e MYSQL_ROOT_PASSWORD=rootroot`：设置MySql数据库的root用户密码为rootroot，实际通过环境变量MYSQL_ROOT_PASSWORD指定

> -`d` 表示在后台运行容器，不会阻塞当前终端

## 修改配置文件
由于我们使用Mysql容器，我们需要对配置文件app.ini进行修改
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

# 这里的密码是创建MySql容器时我们定下的密码
USER = root
PASSWORD = rootroot

#127.0.0.1:3306
#HOST = localhost:3306
HOST = mysql:3306

NAME = blog
TABLE_PREFIX = blog_
```

## 构建镜像
在gin-example项目根目录下执行
```shell
docker build -t gin-blog-docker .
```
> `-t`:指定名称
>
> `. `:构建内容为当前上下文目录,从当前目录中寻找dockerfile文件

## 验证镜像
执行命令
```shell
& docker ps

CONTAINER ID   IMAGE             COMMAND                  CREATED       STATUS       PORTS                               NAMES
89abe529cf28   gin-blog-docker   "./gin-example"          4 hours ago   Up 4 hours   0.0.0.0:8000->8000/tcp              laughing_sammet
88606fa6bacc   mysql             "docker-entrypoint.s…"   5 hours ago   Up 5 hours   0.0.0.0:3306->3306/tcp, 33060/tcp   mysql
```
存在刚刚构建的两个镜像

## 创建并运行容器,将golang容器和MySql容器关联
```shell
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker
```

```shell
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 2024/03/18 06:36:18 Error 1049 (42000): Unknown database 'blog'
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /auth                     --> github.com/kingsill/gin-example/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/kingsill/gin-example/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/kingsill/gin-example/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/kingsill/gin-example/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/kingsill/gin-example/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/kingsill/gin-example/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/kingsill/gin-example/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/kingsill/gin-example/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/kingsill/gin-example/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/kingsill/gin-example/routers/api/v1.DeleteArticle (4 handlers)
2024/03/18 06:58:38 Actual pid is 1
```

> `--link`：可以在容器内直接使用其关联的容器别名进行访问，而不是通过ip
> 但是`--link`只能解决单机容器间的关联
> 在分布式多机的情况下，需要通过别的方式进行连接

当然，现在依旧是由问题的，我们可以看到这一句`2024/03/18 06:36:18 Error 1049 (42000): Unknown database 'blog'`
这说明在**我们现在的MySql容器中没有我们之前创建的数据库和表**

如果只是在当前容器内使用，可以直接在MySql容器内进行表格等建立即可

当然如果想要**数据持久化**，并且不随容器删除而删除，多个容器使用，就需要用到挂载数据卷

## 挂载主机目录
```shell
docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v /home/wang2/docker_rep:/var/lib/mysql mysql
```
>`-v /home/wang2/docker_rep:/var/lib/mysql:`这里多出的一句是将宿主机的目录/home/wang2/docker_rep，挂载到容器的/var/lib/mysql，即mysql默认存储数据的位置，完成数据持久化

当然，也可以选择将我们之前的mysql数据直接复制到这里，这一种方法大家可以自己尝试

之后我们可以直接在容器中再进行数据库内容的初始化即可，[我们第一部分的数据初始化部分](https://blog.csdn.net/kingsill/article/details/135445193#t12)

## 重新运行golang容器
```shell
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker
```

## 验证
访问`http://127.0.0.1:8000/auth?username=test&password=test123456`
得到结果如下：
```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE3MTA3NTU5MzQsImlzcyI6Imdpbi1ibG9nIn0.UO686FMYrBT-pXkkIEcNr7g8l-7kqEpjsd_Gim2bRWE"
    },
    "msg": "ok"
}
```
该项目部署到`docker`成功