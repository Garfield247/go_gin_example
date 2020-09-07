package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	gorm.Model

	Name string `gorm:"type:varchar(10);default:''"json:"name"`
	CreatedBy string `gorm:"type:varchar(100);default:'';"json:"created_by"`
	ModifiedBy string `gorm:"type:varchar(100);default:'';"json:"modified_by"`
	State int `gorm:"type:tinyint(3);default:1;"json:"state"`
}

//根据查询条件查询tag
func GetTags(pageNum int ,pageSize int, maps interface{}) (tags []Tag) {
	//var tags Tag
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

//获取tag的总数（某查询条件下的）
func GetTagTotal(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//判断tag名称是否存在
func ExistTagByName(name string) bool  {
	var tag Tag
	db.Select("id").Where("name = ?",name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?",id).First(&tag)
	if tag.ID > 0{
		return true
	}
	return false
}

//添加Tag
func AddTag(name string,state int ,createdBy string) bool {
	db.Create(&Tag{
		Name:name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?",id).Update(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?",id).Delete(&Tag{})
	return true
}
/*
这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。

gorm所支持的回调方法：

创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind

*/
//设置创建时间
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error{
	scope.SetColumn("CreatedOn",time.Now().Unix())
	return nil
}
//设置修改时间
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error{
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}