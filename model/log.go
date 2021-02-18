package model

import (
	"time"

	"gorm.io/gorm"
)

// Log 日志表
type Log struct {
	ID uint
	Name string
	UserID int
	IP   string
	Remark string
	CreatedAt    time.Time
	// User   User `gorm:"foreignKey:UserID"`
}

// InsertLog 插入日志
func InsertLog(db *gorm.DB, name string, userID int, ip string, remark string)  {
	log := Log{
		Name: name,
		UserID: userID,
		IP: ip,
		Remark: remark, 
	}
	db.Create(&log)
}