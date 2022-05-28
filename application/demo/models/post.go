
// Package models  文章表
// author : http://www.liyang.love
// date : 2022-04-23 19:52
// desc : 文章表
package models

import (
	"github.com/reatang/gostart/pkg/database/mysql"
	"gorm.io/gorm"
)

// Post  文章表。
// 说明:
// 表名:posts
// group: Posts
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Posts
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Posts
// version:2022-04-23 19:52
type Post struct {
	gorm.Model

	// 外键
	User User `gorm:"foreignKey:user_id"`

	// 字段
	Title       string           `gorm:"column:title" json:"title"`              //type:string           comment:字符串      version:2022-04-23 19:52
	UserId      int64            `gorm:"column:user_id" json:"userId"`           //type:int64            comment:数字        version:2022-04-23 19:52
	Content     string           `gorm:"column:content" json:"content"`          //type:string           comment:内容        version:2022-04-23 19:52
}

// TableName 表名:posts，文章表。
// 说明:
func (Post) TableName() string {
	return "posts"
}

type PostDao struct {}

func GetPostDao() *PostDao {
	return &PostDao{}
}

func GetPostModel() *gorm.DB {
	return mysql.GetDB().Model(&Post{})
}

// Create 创建
func (*PostDao) Create(post *Post) *gorm.DB {
	return mysql.GetDB().Create(post)
}

// First 用ID获取数据
func (*PostDao) First(id int64) (post Post, err error) {
	tx := mysql.GetDB().Preload("User").First(&post, id)

	if tx.Error != nil {
		err = tx.Error
	}

	return
}

// Delete 用ID获取数据
func (*PostDao) Delete(id int64) (post Post, err error) {
	tx := mysql.GetDB().Delete(&post, id)

	if tx.Error != nil {
		err = tx.Error
	}

	return
}
