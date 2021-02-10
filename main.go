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
		config.InitAuthConfig()
	}

  dbFile := "test.db"
  db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
	// fmt.Println(db)

  // Migrate the schema
  db.AutoMigrate( &model.User{}, &model.Log{})


	router := gin.Default()
	gin.ForceConsoleColor()
	currentDir,_ := os.Getwd()
	fmt.Println("当前路径：",currentDir)

	router.GET("/ping", func(c *gin.Context) { api.Ping(c) })
	router.POST("/api/login", func(c *gin.Context) { api.Login(c) })
	router.GET("/api/users", func(c *gin.Context) { api.UserIndex(c, db ) })
	// router.POST("/api/version", func(c *gin.Context) { api.UploadVersion(c, conn) })
	// router.POST("/api/del_versions", func(c *gin.Context) { api.DelVersion(c, conn, currentDir) })
	// router.GET("/api/versions", func(c *gin.Context) { api.VersionIndex(c, conn) })
	// router.GET("/api/logs", func(c *gin.Context) { api.LogIndex(c, conn) })
	
	// router.GET("api/v1/version", func(c *gin.Context) { api.CheckVersion(c, conn) })
	// router.POST("api/v1/version", func(c *gin.Context) { api.CheckVersion(c, conn) })
	
	// router.POST("/api/uploads", func(c *gin.Context) { api.UploadFile(c, uploadsPath) })

	router.Static("/dist/", "./dist/")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"/dist")
	})

	router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	router.NoRoute((func(c *gin.Context) { api.NotFind( c) }))

	port := config.GetAuthIni("port")
	port = "0.0.0.0:"+port
	router.Run(port)

  // // Create
  // db.Create(&Product{Code: "D42", Price: 100})

  // // Read
  // var product Product
  // db.First(&product, 1) // find product with integer primary key
  // db.First(&product, "code = ?", "D42") // find product with code D42

  // // Update - update product's price to 200
  // db.Model(&product).Update("Price", 200)
  // // Update - update multiple fields
  // db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // // Delete - delete product
  // db.Delete(&product, 1)
}