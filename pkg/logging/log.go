package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 类型声明，基于int建立level，方便后续维护
type Level int

var (
	//传入写log文件的句柄
	F *os.File

	//默认的前缀
	DefaultPrefix = ""

	//这里的定义在后续的caller中被调用，该参数指定了要跳过的调用堆栈帧数，每个调用堆栈帧代表了代码中的一个函数调用。
	//以info写入log为例，我们这里调用caller的函数为setPrefix，Info调用setPrefix，适用Info进行写入的函数调用info。
	//                                     0                      1                    2
	//而我们想要得到的信息就是调用info的函数的信息。所以我们这里设置的跳过调用堆栈帧数为2
	DefaultCallerDepth = 2

	//提前定义logger记录器，方便维护阅读
	logger *log.Logger

	logPrefix = ""

	//结合level的定义，方便读取维护
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// 实现枚举
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup 自定义logger的初始化
func Setup() {
	var err error

	//得到log文件句柄
	F, err = openLogFile()
	if err != nil {
		log.Fatalln(err)
	}

	//创建一个新的日志记录器
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info 这里先设置每条log的前缀部分，首先为log模式，这里为info；然后为具体到某个函数第几行出错；接下来为时间；最后为日志信息
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// 从进程中读取当前运行的函数信息
func setPrefix(level Level) {
	//获取文件名，具体行数，是否读取成功
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok { //获取成功
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else { //获取失败，则前缀不加如具体的文件名和行号
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	//将前缀写入log文件
	logger.SetPrefix(logPrefix)
}
