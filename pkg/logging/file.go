package logging

import (
	"fmt"
	"github.com/kingsill/gin-example/pkg/file"
	"github.com/kingsill/gin-example/pkg/setting"
	"os"
	"time"
)

// 返回log文件的前缀路径，算是一个具有仪式感的函数
func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

// 获得log文件的整体路径，以当前日期作为.log文件的名字 runtime/log20010212.log
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// 打开日志文件，返回写入的句柄handle
func openLogFile() (*os.File, error) {

	//获取文件整体路径
	fileName := getLogFileFullPath()

	//创建目录
	mkDir()

	//如果.log文件不存在，这里会创建一个
	handle, err := file.Open(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open:%s\n", fileName)
	}

	return handle, nil
}

// 创建log目录
func mkDir() {
	//获得当前目录 dir: /home/wang2/gin-example
	dir, _ := os.Getwd()

	//检查目录访问权限
	perm := file.CheckPermission(getLogFilePath())
	if perm == true {
		panic("Permission denied")
	}

	//如果目录不存在，创建目录
	err := file.IsNotExistMkDir(dir + "/" + getLogFilePath())
	if err != nil {
		panic(err)
	}
}
