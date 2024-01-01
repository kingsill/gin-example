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

// GetTagTotal 查询tag总数
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

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
