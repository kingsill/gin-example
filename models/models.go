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
