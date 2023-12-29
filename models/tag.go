package models

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
