package model

import (
	"time"

	"gorm.io/gorm"
)

// File 用户表
type File struct {
	ID           uint
	Name string
	Link string
	CreatedAt    time.Time
}


// InsertFile 插入日志
func InsertFile(db *gorm.DB, Link string, name string)  {
	log := File{
		Name: name,
		Link: Link,
	}
	db.Create(&log)
}