package model

import "gorm.io/gorm"

// Log 日志表
type Log struct {
	gorm.Model
	ID uint32
	Name string
	UserID int
	IP   string
	Remark string
	User User
}