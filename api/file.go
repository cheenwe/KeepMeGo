package api

import (
	"KeepMeGo/model"
	"KeepMeGo/pkg/upload"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UploadFile 处理文件上传
func UploadFile(c *gin.Context, db *gorm.DB, uploadsPath string)  {
	var contentLength int64
	contentLength = c.Request.ContentLength
	if contentLength<=0 || contentLength>1024*1024*1024*5{
		log.Printf("contentLength error\n")
		return
	}
	contentType,hasKey := c.Request.Header["Content-Type"]
	if  !hasKey{
		log.Printf("Content-Type error\n")
		return
	}
	if len(contentType)!=1{
		log.Printf("Content-Type count error\n")
		return
	}
	contentTypeVue := contentType[0]
	const BOUNDARY string = "; boundary="
	loc := strings.Index(contentTypeVue, BOUNDARY)
	if -1==loc{
		log.Printf("Content-Type error, no boundary\n")
		return
	}
	boundary := []byte(contentTypeVue[(loc+len(BOUNDARY)):])
	log.Printf("[%s]\n\n", boundary)
	//
	readData := make([]byte, 1024*12)
	var readTotal int = 0
	for {
		fileHeader, fileData, err := upload.ParseFromHead(readData, readTotal, append(boundary, []byte("\r\n")...), c.Request.Body)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		fileFolderPath := "/dist/uploads/"
		// 存储收到的文件 
		fileName := uploadsPath+fileHeader.FileName
		filePath :=  fileFolderPath+fileHeader.FileName
		log.Printf("file :%s\n", fileName)

		f, err := os.Create(fileName)
		if err != nil {
			log.Printf("create file fail:%v\n", err)
			return
		}
		f.Write(fileData)
		fileData = nil
		//需要反复搜索boundary
		tempData, reachEnd, err := upload.ReadToBoundary(boundary, c.Request.Body, f)
		f.Close()
 

		model.InsertFile(db, filePath, "")

		// 插入日志  开始
		cip := c.ClientIP()
		cmodel := "文件"
		cuserID := 0 //TODO 获取用户ID
		cremark := "上传文件: "+filePath
		model.InsertLog(db, cmodel, cuserID, cip, cremark)  
		// 插入日志  结束

		if err != nil {
			log.Printf("%v\n", err)
			return
		}
		if reachEnd{
			//c.JSON(200, result)
			c.JSON(200, gin.H{
				"msg": "上传成功！",
				"code": 1,
				"data": filePath,
			})
		} else {
			copy(readData[0:], tempData)
			readTotal = len(tempData)
			continue
		}
	}
}


// FileIndex 用户列表
func  FileIndex(c *gin.Context, db *gorm.DB) {
	// 插入日志  开始
	cip := c.ClientIP()
	cmodel := "用户"
	cuserID := 0 //TODO 获取用户ID
	cremark := "访问文件列表"
	model.InsertLog(db, cmodel, cuserID, cip, cremark)  
	// 插入日志  结束

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
		db = db.Unscoped().Model(model.File{}) //查询对应的数据库表
	}else{
		db = db.Model(model.File{}) //查询对应的数据库表
	}

	var files []model.File
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
	if err := db.Order("id DESC").Offset(offset).Limit(perPage).Find(&files).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 500,
			"success":    0,
			"message" : "查询数据异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"success":    1,
		"data" : files,
		"total": total,
		"page" : page,
		"per_page": perPage,
	})
	return 
}