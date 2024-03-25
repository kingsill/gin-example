package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kingsill/gin-example/pkg/setting"
)

// 定义一个全局的数据库连接变量
var db *gorm.DB

// Model 设定常用结构体，可以作为匿名结构体嵌入到别的表格对应的结构体
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func Setup() {
	var err error

	//使用gorm框架初始化数据库连接
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	//自定义默认表的表名，使用匿名函数，在原默认表名的前面加上配置文件中定义的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	//gorm默认使用复数映射，当前设置后即进行严格匹配
	db.SingularTable(true)
	//log记录打开
	db.LogMode(true)

	//进行连接池设置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//替换Create和Update回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	//添加删除的回调CallBacks
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

/*// 将以下定义为init函数
func init() {
	//var (
	//	err                                               error
	//	dbType, dbName, user, password, host, tablePrefix string
	//)

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

	//替换Create和Update回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	//添加删除的回调CallBacks
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}*/

// CloseDB 与数据库断开连接函数
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback 在创建记录时设置 `CreatedOn`, `ModifiedOn`
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	//检查作用域是否出错
	if scope.HasError() {
		return
	}

	//将当前时间记录为时间戳
	nowTime := time.Now().Unix()

	//检查'CreatedOn'字段是否存在，如果存在并且为空，则填充
	if createTimeField, ok := scope.FieldByName("CreatedOn"); ok && createTimeField.IsBlank {
		createTimeField.Set(nowTime)
	}

	//同上
	if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok && modifyTimeField.IsBlank {
		modifyTimeField.Set(nowTime)
	}
}

// updateTimeStampForUpdateCallback 在更新记录时设置 `ModifyOn`
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {

	//检查是否指定特定的更新行，当指定更新字段时不修改'ModifiedOn'
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 设定delete操作的callback逻辑
func deleteCallback(scope *gorm.Scope) {
	//检查作用域
	if !scope.HasError() {

		var extraOption string

		//尝试获取delete_option，即操作过过程的其他附加选项，如果存在的话
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		//根据字段名查找模型中的字段，返回其相关信息及是否存在
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		//unscoped代表不忽略软删除，如果不忽略软删除和存在deletedOn字段
		if !scope.Search.Unscoped && hasDeletedOnField {

			//创建原生的sql语句
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),                            //表名
				scope.Quote(deletedOnField.DBName),                 //转义字段名
				scope.AddToVars(time.Now().Unix()),                 //时间戳添加，这里使用AddToVars方法作为变量插入是为了保证安全性
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //CombinedConditionSql()用于获取当前查询条件的sql表示式，使用此避免出现原本sql语句应有功能的确实
				addExtraSpaceIfExist(extraOption),                  //之后再额外添加附加选项
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf( //内容同上进行理解
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// 判断是否为空来进行空格插入，防止sql注入，保证安全性
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
