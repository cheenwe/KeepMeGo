package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	ID uint32
	Name string
	Email   string
	Password string
	Token string
}