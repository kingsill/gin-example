package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 适用枚举，将所有固定的量提前列出在这里，方便后期维护	我们这里将原来的var修改为const
const (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// 返回log文件的前缀路径，算是一个具有仪式感的函数
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// 获得log文件的整体路径，以当前日期作为.log文件的名字
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// 打开日志文件，返回写入的句柄handle
func openLogFile(filePath string) *os.File {
	//根据文件目录是否存在进行判断
	_, err := os.Stat(filePath)
	switch {
	//目录不存在
	case os.IsNotExist(err):
		mkDir()

	//权限不够
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	//如果.log文件不存在，这里会创建一个
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

// 创建log目录
func mkDir() {
	//获得当前目录 dir: /home/wang2/gin-example
	dir, _ := os.Getwd()

	//适用MKdirAll会直接创建所有依赖的父目录，减少报错的可能性
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
