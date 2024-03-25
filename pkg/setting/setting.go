package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 这里对应ini文件中中的app部分,注意保持同样的大驼峰写法
type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// AppSetting 由于MapTo视线中约束适用指针,我们地址入参
var AppSetting = &App{}

// Server Server和Database与App类似,这里不再解释
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

// Setup
func Setup() {
	//配置文件加载
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//将app section 部分映射到AppSetting结构体上
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	//将图片最大大小设置从5字节Byte转换为5兆字节MB
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	//将读取时自动转换的类型转换为时间间隔了，只不过是最小单位纳秒
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}

// -----------------------     优化配置前结构              ---------------------------------------------------------
/*var (
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
*/
