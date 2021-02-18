package api

import (
	"KeepMeGo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LogIndex 日志列表
func  LogIndex(c *gin.Context, db *gorm.DB) {
	var total int64
	pageNum := "1"
	perPageNum := "10"
	if c.Query("page") != "" {
		pageNum = c.Query("page")
	}
	if c.Query("per_page") != "" {
		perPageNum = c.Query("per_page")
	}

	page,_          := strconv.Atoi(pageNum)
	perPage,_  := strconv.Atoi(perPageNum)
	//此处用了PostForm的请求方法
	del := c.Query("del") //查询删除过的记录

	if del == "1" {
		db = db.Unscoped().Model(model.Log{}) //查询对应的数据库表
	}else{
		db = db.Model(model.Log{}) //查询对应的数据库表
	}

	var logs []model.Log
	
	if err := db.Count(&total).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"success":    0,
			"message" : "查询数据异常",
		})
		return
	}
	//此时的total是查询到的总数
	offset := (page-1)*perPage
	if err := db.Order("id DESC").Offset(offset).Limit(perPage).Find(&logs).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"success":    0,
			"message" : "查询数据异常",
		})
		return
	}

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "日志"
	cuserID := 0 //TODO 获取用户ID
	cremark := "访问日志列表"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"success":    1,
		"data" : logs,
		"total": total,
		"page" : page,
		"per_page": perPage,
	})
	return 
}
