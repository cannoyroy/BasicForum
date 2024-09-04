package database

import (
	"fmt"
	"log"
	"server/internal/global"
	"server/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	user := global.Config.GetString("database.user")
	pass := global.Config.GetString("database.password")
	host := global.Config.GetString("database.host")
	port := global.Config.GetString("database.port")
	name := global.Config.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connec222ting to database: %s", err)
	}

	// 迁移框架
	db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Reports{}, &models.Trash{})
	return db
}
