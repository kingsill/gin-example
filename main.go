package main

import (
	"fmt"
	"net/http"

	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/routers"
)

func main() {

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
}
