// 自动生成模板Student
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Student struct {
      gorm.Model
      Name  string `json:"name" form:"name" gorm:"column:name;comment:''"`
      Sex  *bool `json:"sex" form:"sex" gorm:"column:sex;comment:''"`
      Age  int `json:"age" form:"age" gorm:"column:age;comment:''"` 
}


func (Student) TableName() string {
  return "student"
}
