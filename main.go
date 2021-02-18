package main

import (
	"KeepMeGo/api"
	"KeepMeGo/model"
	"KeepMeGo/pkg/config"
	"KeepMeGo/pkg/util"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	configFile := "./config.ini"

	authFileExist := util.FileExist(configFile)
	if !authFileExist {
		f, _ :=os.Create(configFile)
		f.Close()
		config.InitConfigFile()
	}

	dbFile := config.GetConfigIni("config", "db")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate( &model.User{}, &model.Log{}, &model.File{})

	router := gin.Default()
	gin.ForceConsoleColor()
	currentDir,_ := os.Getwd()
	fmt.Println("当前路径：",currentDir)

	fileFolderPath := "/dist/uploads/"

	uploadsPath := currentDir+fileFolderPath
	os.MkdirAll(uploadsPath, os.ModePerm)

	router.GET("/ping", func(c *gin.Context) { api.Ping(c) })
	router.GET("/api/logs", func(c *gin.Context) { api.LogIndex(c, db) })
	router.POST("/api/login", func(c *gin.Context) { api.Login(c, db) })

	router.GET("/api/users", func(c *gin.Context) { api.UserIndex(c, db) })
	router.POST("/api/users", func(c *gin.Context) { api.UserCreate(c, db) })
	router.GET("/api/users/:id", func(c *gin.Context) { api.UserRead(c, db) })
	router.PUT("/api/users/:id", func(c *gin.Context) { api.UserUpdate(c, db) })
	router.DELETE("/api/users/:id", func(c *gin.Context) { api.UserDelete(c, db) })

	router.POST("/api/uploads", func(c *gin.Context) { api.UploadFile(c, db, uploadsPath) })
	router.GET("/api/uploads", func(c *gin.Context) { api.FileIndex(c, db) })

	router.Static("/dist/", "./dist/")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"/dist")
	})

	router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	router.NoRoute((func(c *gin.Context) { api.NotFind( c, db) }))

	port := config.GetConfigIni("config", "port")
	port = "0.0.0.0:"+port
	router.Run(port)
}