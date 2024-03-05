package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/routers"
	"log"
	"syscall"
)

// @title gin-example
// @version 1.0

func main() {

	/*	之前的版本
		//将原来创建的默认router和建立对应的handler绑定
			router := routers.InitRouter()

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
	*/

	//当前版本，建议对比源代码进行理解
	//进行配置的导入，在setting包中进行设置

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	//newsever返回一个初始化的endlessServer对象
	server := endless.NewServer(endPoint, routers.InitRouter())
	//输出当前进程的PID
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	//启动服务
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
