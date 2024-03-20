# 前情提要以及需求产生
[第七部分我们实现了软删除](https://blog.csdn.net/kingsill/article/details/136850334?spm=1001.2014.3001.5501)，那么我们**什么时候硬删除**呢？

一般有**两种解决方案**：
- 另外一套硬删除接口
- 定时任务清理无效数据 --我们这里使用第二种进行解决

# 实现
## 安装cron包
```shell
go get -u github.com/robfig/cron
```
## 编写硬删除代码
打开 models 目录下的 tag.go、article.go 文件，分别添加以下代码
1. `tag.go`
```go
func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
```
2. `article.go`
```go
func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}

```
## 编写 Cron
在 项目根目录下新建 `cron.go` 文件，用于编写定时任务的代码，写入文件内容
```go
package main

import (
	"time"
	"log"

	"github.com/robfig/cron"

	"github.com/EDDYCJY/go-gin-example/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}

```

