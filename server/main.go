package main

import (
	"log"
	"server/internal/pkg/database"
	"server/internal/router"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.Init()
	services.ServiceInit(db)
	r := gin.Default()
	router.Init(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
