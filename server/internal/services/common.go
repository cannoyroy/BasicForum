package services

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ServiceInit(dbb *gorm.DB) {
	db = dbb
}
