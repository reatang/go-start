// Package models  用户
// author : http://www.liyang.love
// date : 2022-04-23 15:39
// desc : 用户
package models

import (
	"github.com/reatang/gostart/pkg/database/mysql"
	"gorm.io/gorm"
	"time"
)

// User  用户。
// 说明:
// 表名:users
// group: Users
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Users
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Users
// version:2022-04-23 15:39
type User struct {
	Id         int64        `gorm:"column:id;primaryKey" json:"Id"`      //type:int64        comment:主键    version:2022-04-23 15:39
	Name       string       `gorm:"column:name" json:"Name"`             //type:string       comment:名称    version:2022-04-23 15:39
	Age        int          `gorm:"column:age" json:"Age"`               //type:int          comment:年龄    version:2022-04-23 15:39
	Birthday   time.Time   `gorm:"column:birthday" json:"Birthday"`     //type:*time.Time   comment:生日    version:2022-04-23 15:39
}

// TableName 表名:users，用户。
// 说明:
func (User) TableName() string {
	return "users"
}

type UserDao struct {}

func GetUserDao() *UserDao {
	return &UserDao{}
}

// Create 创建
func (*UserDao) Create(user *User) *gorm.DB {
	return mysql.GetDB().Create(user)
}

// First 用ID获取数据
func (*UserDao) First(id int64) (user *User) {
	mysql.GetDB().First(&user, id)

	return
}

