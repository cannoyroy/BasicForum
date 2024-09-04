package router

import (
	"server/internal/midwares"
	
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(midwares.Cors())

	r.POST("/api/user/login", func(c *gin.Context) {
		services.Login(c)
	})

	r.POST("api/user/reg", func(c *gin.Context) {
		services.Reg(c)
	})

	r.POST("api/student/post", func(c *gin.Context) {
		services.Post(c)
	})

	r.GET("api/student/post", func(c *gin.Context) {
		services.GetPost(c)
	})

	r.PUT("api/student/post", func(c *gin.Context) {
		services.PutPost(c)
	})

	r.DELETE("api/student/post", func(c *gin.Context) {
		services.DelePost(c)
	})

	r.POST("api/student/report-post", func(c *gin.Context) {
		services.ReportPost(c)
	})

	r.GET("api/student/report-post", func(c *gin.Context) {
		services.GetReport(c)
	})

	r.GET("/api/admin/report", func(c *gin.Context) {
		services.AdminGetReport(c)
	})

	r.POST("api/admin/report", func(c *gin.Context) {
		services.AdminProReport(c)
	})
}
