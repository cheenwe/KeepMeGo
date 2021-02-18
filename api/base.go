package api

import (
	"KeepMeGo/model"
	"KeepMeGo/pkg/config"
	"KeepMeGo/pkg/util"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login 登录接口
func  Login(c *gin.Context, db *gorm.DB) {
	result := map[string]interface{}{}

	result["msg"] = "参数错误"
	result["code"] = 0

	log.Printf("req Username: %v", c.PostForm("name"))
	log.Printf("req password: %v", c.PostForm("password"))

	reqUsername := c.PostForm("name") //从表单中查询参数
	reqPassword := c.PostForm("password")//从表单中查询参数
		
	var resMsg string
	resMsg = "1"
	log.Printf("token: %v", resMsg)
    if reqUsername != "" {
        log.Println("====== Bind By Query String ======")

		username := config.GetAuthIni("username")
		password := config.GetAuthIni("password")
		log.Printf("username: %v", username)
		log.Printf("password: %v", password)

		token := config.GetAuthIni("token")
		log.Printf("token: %v", token)

		if reqUsername == username && reqPassword == password {
			result["success"] = 1
			result["token"] = token
			result["code"] = 200
			result["msg"] = "登录成功"
			result["time"] =  time.Now().Format("2006-01-02 15:04:05")
			resMsg = "成功"
		} else {
			result["msg"] = "用户名或密码错误"
			result["success"] = 0
			result["code"] = 200
			result["data"] = "用户名或密码错误"
			result["time"] =  time.Now().Format("2006-01-02 15:04:05")
			resMsg = "失败"
		}
    }

	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "登录系统"
	cuserID := 0 //TODO 获取用户ID
	cremark := "用户【"+reqUsername+"】登录"+ resMsg
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束
	c.JSON(200, result)
}

//Ping Ping接口
func Ping(c *gin.Context) {
	result := map[string]interface{}{}
	now := time.Now().Format("2006-01-02 15:04:05")
	result["msg"] = "success!"
	result["success"] = 1
	result["code"] = 200
	result["data"] = util.RandomString(12)
	result["time"] = now
	c.JSON(200, result)
}


//NotFind 请求不存在
func NotFind(c *gin.Context, db *gorm.DB) {
	result := map[string]interface{}{}
	result["msg"] = "Not Find!"
	result["success"] = 0
	result["code"] = 404
	result["data"] = "error!"
	url := c.Request.RequestURI
	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "访问系统"
	cuserID := -1 //TODO 获取用户ID
	cremark := "访问失败: "+url
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束
	c.JSON(404, result)
}