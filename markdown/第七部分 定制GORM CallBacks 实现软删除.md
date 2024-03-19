# 涉及知识点
`GORM` 本身是由回调驱动的，所以我们可以根据需要完全定制`gorm`
> GORM itself is powered by Callbacks, so you could fully customize GORM as you want

gorm包含以下四类`CallBacks`：
- 注册一个新的回调
- 删除现有的回调
- 替换现有的回调
- 注册回调的顺序
本节我们使用 `替换现有的回调 `
# 本节目标

## 项目原有问题
我们可以发现在`article.go、tag.go`中都有涉及到以下方法：
```go
// BeforeCreate 与tag的逻辑相同 为了插入createOn时间戳
func (...) BeforeCreate(...) error {
...
}

func (...) BeforeUpdate(...) error {
...
}
```
这是为了更新表中的`createdon`和`ModifiedOn`，这里有两个涉及以下字段的表，我们在两个`model`中都设置这两个方法，为了避免出现更多的表而再次**编写更多次**类似方法，我们通过`callbacks`来实现该功能。

# 实现CallBacks
## 新增方法
编辑`models`目录下的`models.go`,新增以下两个方法
```go
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

```


## 注册CallBacks
上面我们定义好回调方法，这一部分我们将其**注册到GORM的钩子**，替换其原本的`Create`和`Update`回调

在`models.go`中的`init`函数中，增加以下语句：
```go
db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
```

## 验证
使用新建文章或包含这两个字段的表格，可以看到这两个字段功能都正常

# 通过`callbacks`实现软硬删除
当然，如果您的模型包含了一个 `gorm.deletedat `字段（`gorm.Model `已经包含了该字段)，它将**自动获得软删除的能力**！

我们这里是尝试通过`callbacks`实现

## 实现callbacks
1. 在`model struct`增加`delectedon`变量

```go
// Model 设定常用结构体，可以作为匿名结构体嵌入到别的表格对应的结构体
type Model struct {
ID         int `gorm:"primary_key" json:"id"`
CreatedOn  int `json:"created_on"`
ModifiedOn int `json:"modified_on"`
DeletedOn  int `json:"deleted_on"`
}
```
这里我们就与gorm自带的gorm.model一致了，最后的deletedon就是为了软删除设置的

2. 打开 models 目录下的 models.go 文件，实现以下方法：
```go
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
scope.Raw(fmt.Sprintf(//内容同上进行理解
"DELETE FROM %v%v%v",
scope.QuotedTableName(),
addExtraSpaceIfExist(scope.CombinedConditionSql()),
addExtraSpaceIfExist(extraOption),
)).Exec()
}
}
}

//判断是否为空来进行空格插入，防止sql注入，保证安全性
func addExtraSpaceIfExist(str string) string {
if str != "" {
return " " + str
}
return ""
}
```

## 注册CallBacks
在 `models.go` 的 `init` 函数中，增加以下删除的回调
```go
db.Callback().Delete().Replace("gorm:delete", deleteCallback)
```

## 验证
到这里，大家可以进行自行进行删除操作，查看deletedOn字段是否存在删除时的时间戳
