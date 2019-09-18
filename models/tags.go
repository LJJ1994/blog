package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	_  = scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagsTotal(maps interface{}) (count int) {
	db.Where(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistsTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name= ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistById(id int) bool {
	var tag Tag
	db.Select("id").Where("id= ?", id).Find(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func EditTag(id int, maps interface{}) bool {
	db.Model(&Tag{}).Where("id= ?", id).Update(maps)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id=? ", id).Delete(&Tag{})
	return true
}

func AddTags(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:       name,
		CreatedBy:   createBy,
		State:      state,
	})
	return true
}