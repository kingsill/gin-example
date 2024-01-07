package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Tag 定义tag表的相关表头
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//以下两个函数都为命名返回，函数内部有相关定义

// GetTags page-size,每页显示的tag数	pageNum，即为从第几条记录开始显示，从0开始计数
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {

	// where查询使用map条件
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags) //db在models.go中已经定义，同一个包可以直接调用
	return
}

// GetTagTotal 查询tag总数，使用时maps会传递相关参数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName 根据name查询tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	//查询词条
	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// AddTag 创建新tag
func AddTag(name string, state int, createdBy string) bool {
	//在对应数据表中创建词条
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// BeforeCreate 建立hook钩子函数，在创建之前插入时间值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	//我们在定义的时候createon是int类型，因此我们这里使用unix方法将时间转变为时间戳
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate 同为hook钩子函数，更新的时候加入修改时间值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// ExistTagByID 根据id查询表中tag是否存在
func ExistTagByID(id int) bool {
	var tag Tag //实例化tag

	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// DeleteTag 根据id删除表中tag
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

// EditTag 修改tag
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
