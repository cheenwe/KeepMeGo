package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	ID           uint
	Name string
	Email   string
	Password string
	Token string 
	CreatedAt    time.Time
  	UpdatedAt    time.Time
}


// // UserSimple 用户表
// type UserSimple struct {
// 	ID           uint
// 	Name string
// }

// // TableName 自定义表名
// func (UserSimple) TableName() string {
//     return "users"
// }