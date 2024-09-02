package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"verse1.0/app/midwares"
	"verse1.0/app/models"
	"verse1.0/app/services"
)

func main() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 初始化数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// 迁移框架
	db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Reports{})

	// 初始化GIN
	// 负责创建并返回一个已经配置好默认中间件的 gin.Engine 实例。
	r := gin.Default()

	r.Use(midwares.Cors())

	r.POST("/api/user/login", func(c *gin.Context) {
		services.Login(c, db)
	})

	r.POST("api/user/reg", func(c *gin.Context) {
		services.Reg(c, db)
	})

	r.POST("api/student/post", func(c *gin.Context) {
		services.Post(c, db)
	})

	r.GET("api/student/post", func(c *gin.Context) {
		services.GetPost(c, db)
	})

	r.PUT("api/student/post", func(c *gin.Context) {
		services.PutPost(c, db)
	})

	r.DELETE("api/student/post", func(c *gin.Context) {
		services.DelePost(c, db)
	})

	r.POST("api/student/report-post", func(c *gin.Context) {
		services.ReportPost(c, db)
	})

	r.GET("api/student/report-post", func(c *gin.Context) {
		services.GetReport(c, db)
	})

	r.GET("/api/admin/report", func(c *gin.Context) {
		services.AdminGetReport(c, db)
	})

	r.POST("api/admin/report", func(c *gin.Context) {
		services.AdminProReport(c, db)
	})

	// 开启服务器
	r.Run(":8080")
}
