package api

import (
	"KeepMeGo/model"
	"KeepMeGo/pkg/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserIndex 用户列表
func  UserIndex(c *gin.Context, db *gorm.DB) {
	
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
		db = db.Unscoped().Model(model.User{}) //查询对应的数据库表
	}else{
		db = db.Model(model.User{}) //查询对应的数据库表
	}

	var users []model.User
	
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
	if err := db.Order("id DESC").Offset(offset).Limit(perPage).Find(&users).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"success":    0,
			"message" : "查询数据异常",
		})
		return
	}

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "访问用户列表"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"success":    1,
		"data" : users,
		"total": total,
		"page" : page,
		"per_page": perPage,
	})
	return 
}

// UserCreate 用户创建
func  UserCreate(c *gin.Context, db *gorm.DB) {

	name := c.PostForm("name") //从表单中查询参数
	email := c.PostForm("email") //从表单中查询参数
	password := c.PostForm("password")//从表单中查询参数

	user := model.User{
		Name: name, 
		Email: email,
		Password: password,
		Token: util.RandomString(12)}

	result := db.Create(&user) // 通过数据的指针来创建
	if err := result.Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"success":    0,
			"message" : "用户创建失败",
		})
		return
	}

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "创建用户【"+name+"】"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success":    1,
		"message": "success",
		"data" : user.ID,
	})
	return 
}


// UserRead 用户读取
func  UserRead(c *gin.Context, db *gorm.DB) {
	var user model.User
	id := c.Param("id")//查询参数 id
	del := c.Query("del") //查询删除过的记录

	if del == "1" {
		db.Unscoped().First(&user, id) // 根据整形主键查找
	}else{
		db.First(&user, id) // 根据整形主键查找
	}

	if user.ID == 0{
		c.JSON(http.StatusOK, gin.H{
			"code" : 404,
			"success":  0,
			"message" : "用户不存在",
		})
		return
	}

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "访问用户【"+id+"】"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success":    1,
		"message": "success",
		"data" : user,
	})
	return 
}
    
// UserUpdate 用户更新
func  UserUpdate(c *gin.Context, db *gorm.DB) {
	var user model.User
	id := c.Param("id") //查询参数 id
	// db.First(&user, "id = ?", id) // 查询id
	db.First(&user, id) // 根据整形主键查找

	name := c.PostForm("name") //从表单中查询参数
	email := c.PostForm("email") //从表单中查询参数
	// password := c.PostForm("password")//从表单中查询参数

	db.Model(&user).Updates(model.User{Name: name, Email: email, UpdatedAt: time.Now(), Token: util.RandomString(12) })

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "更新用户【"+id+"】"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success":    1,
		"message": "success",
		"data" : user,
	})
	return 
}

// UserDelete 用户删除
func  UserDelete(c *gin.Context, db *gorm.DB) {
	// var user model.User
	id := c.Param("id") //查询参数 id
	del := c.Query("del") //查询删除过的记录

	if del == "1" {
		db.Debug().Unscoped().Where("id = ?", id).Delete(&model.User{})
	}else{
		db.Where("id = ?", id).Delete(&model.User{})
	}

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "删除用户【"+id+"】"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success":    1,
		"message": "success",
		"data" : id,
	})
	return 
}