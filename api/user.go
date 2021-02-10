package api

import (
	"KeepMeGo/model"
	"KeepMeGo/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserIndex 用户列表
func  UserIndex(c *gin.Context, db *gorm.DB) {
	var total int64
	page,_          := strconv.Atoi(c.DefaultPostForm("page","1"))
	perPage,_  := strconv.Atoi(c.DefaultPostForm("perPage","10"))
	//此处用了PostForm的请求方法
	db = db.Model(model.User{}) //查询对应的数据库表

	var users []model.User
	//这里的models是对数据库进行初始化以及Gorm中的model结构体定义，如下：
	/* 
	var db *gorm.DB
	type Model struct {
		ID         int        `gorm:"primary_key" json:"id"`
		CreatedOn  int        `json:"-"`
		ModifiedOn int        `json:"-"`
		DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	}*/
	if err := db.Count(&total).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"message" : "查询数据异常",
		})
		return
	}
	//此时的total是查询到的总数
	offset := (page-1)*perPage
	if err := db.Order("id DESC").Offset(offset).Limit(perPage).Find(&users).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"message" : "查询数据异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data" : users,
		"total": total,
		"page" : page,
		"per_page": perPage,
	})
	return 
}

// UserCreate 用户创建
func  UserCreate(c *gin.Context, db *gorm.DB) {
	var user model.User

	name := c.PostForm("name") //从表单中查询参数
	email := c.PostForm("email") //从表单中查询参数
	password := c.PostForm("password")//从表单中查询参数

	db.Create(&model.User{Name: name, Email: email, Password: password, Token: util.RandomString(12)})

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data" : user,
	})
	return 
}

  
// UserRead 用户读取
func  UserRead(c *gin.Context, db *gorm.DB) {
	var user model.User
	id := c.Query("id") //查询参数 id
	db.First(&user, id) // 根据整形主键查找

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data" : user,
	})
	return 
}

    
// UserUpdate 用户更新
func  UserUpdate(c *gin.Context, db *gorm.DB) {
	var user model.User
	id := c.Query("id") //查询参数 id
	db.First(&user, id) // 根据整形主键查找

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data" : user,
	})
	return 
}  

// UserDelete 用户删除
func  UserDelete(c *gin.Context, db *gorm.DB) {
	var user model.User
	id := c.Query("id") //查询参数 id
	db.First(&user, id) // 根据整形主键查找

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data" : user,
	})
	return 
}

  


