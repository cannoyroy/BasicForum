package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/internal/models"
	utils "server/pkg"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var user models.Users
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{"code": 200506, "data": nil, "msg": "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusOK, gin.H{"code": 200507, "data": nil, "msg": "密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user_id": user.UserID, "username": user.Username, "name": user.Name, "user_type": user.UserType}, "msg": "success"})
}

func Reg(c *gin.Context) {
	var req models.RegRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	// username只能由数字组成
	if !(utils.IsNumber(req.Username)) {
		c.JSON(http.StatusOK, gin.H{"code": 200502, "data": nil, "msg": "用户名必须为纯数字"})
		return
	}
	// 密码长度大于8位小于16位
	if len(req.Password) <= 8 || len(req.Password) >= 16 {
		c.JSON(http.StatusOK, gin.H{"code": 200503, "data": nil, "msg": "密码长度必须在8-16位"})
		return
	}
	// 用户类型只有1和2
	if req.UserType != 1 && req.UserType != 2 {
		c.JSON(http.StatusOK, gin.H{"code": 200504, "data": nil, "msg": "用户类型错误"})
		return
	}
	// 用户不能存在
	var user models.Users
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 继续注册
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
	} else {
		// 用户已存在
		c.JSON(http.StatusOK, gin.H{"code": 200505, "data": nil, "msg": "用户名已存在"})
		return
	}

	newUser := models.Users{
		UserID:    utils.GenerateID(req.Username, time.Now()),
		Username:  req.Username,
		Name:      req.Name,
		Password:  req.Password,
		UserType:  req.UserType,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&newUser).Error; err != nil {
		// println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func Post(c *gin.Context) {
	var req models.PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var newPost models.Posts

	// 查找用户信息
	var nowUser models.Users
	if err := db.Where("user_id = ?", req.UserID).First(&nowUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	} else {
		newPost.UserID = req.UserID
		newPost.Username = nowUser.Username
		newPost.Name = nowUser.Name
		newPost.UserType = nowUser.UserType
	}

	// 找一下当前的post_id
	var check_ex = true
	var lastPost models.Posts
	if err := db.Last(&lastPost).Error; err != nil {
		check_ex = false
	}
	// 赋值id
	if check_ex {
		newPost.PostID = lastPost.PostID + 1
	} else {
		newPost.PostID = 1
	}

	newPost.Content = req.Content
	newPost.Reason = "/"
	newPost.State = -1
	newPost.CreatedAt = time.Now()

	if err := db.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"PostID": newPost.PostID}, "msg": "success"})
}

func GetPost(c *gin.Context) {
	var PostAll []models.Posts
	if err := db.Where("state != ?", 1).Find(&PostAll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	type out struct {
		UserID    int
		Content   string
		CreatedAt time.Time
		PostID    int
	}
	var OUTposts []out
	for _, v := range PostAll {
		var temp out
		temp.UserID = v.UserID
		temp.Content = v.Content
		temp.CreatedAt = v.CreatedAt
		temp.PostID = v.PostID
		OUTposts = append(OUTposts, temp)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": OUTposts}, "msg": "success"})
}

func PutPost(c *gin.Context) {
	var req models.RePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarPost models.Posts
	if err := db.Where("post_id = ? AND user_id = ?", req.PostID, req.UserID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	TarPost.Content = req.Content
	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func DelePost(c *gin.Context) {
	// var req DePostRequest
	UserID := c.DefaultQuery("user_id", "0")
	PostID := c.DefaultQuery("post_id", "0")

	var TarPost models.Posts
	if err := db.Where("post_id = ? AND user_id = ?", PostID, UserID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "不存在该帖子或者该帖子不属于该用户"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	if err := db.Delete(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ReportPost(c *gin.Context) {
	var req models.ReportPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarPost models.Posts
	if err := db.Where("post_id = ? AND state != ?", req.PostID, 1).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	TarPost.Reason = req.Reason
	TarPost.State = 0

	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	RepotFresh := models.Reports{
		PostID:       req.PostID,
		ReportUserID: req.UserID,
		Reason:       req.Reason,
		CreatedAt:    time.Now(),
	}
	if err := db.Create(&RepotFresh).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func GetReport(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")
	// println(UserID)

	var TarReport []models.Reports
	if err := db.Where("report_user_id = ?", UserID).Find(&TarReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	type out struct {
		PostID  int
		Content string
		Reason  string
		State   int
	}
	var OUTposts []out
	for _, v := range TarReport {
		var temp out
		temp.PostID = v.PostID
		temp.Reason = v.Reason

		var RawPost models.Posts
		if err := db.Where("post_id = ?", v.PostID).Find(&RawPost).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
		temp.Content = RawPost.Content
		temp.State = RawPost.State

		OUTposts = append(OUTposts, temp)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"report_list": OUTposts}, "msg": "success"})
}

func AdminGetReport(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")
	UserID = UserID
	var TarReport []models.Reports
	if err := db.Find(&TarReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	type out struct {
		Username string
		Content  string
		Reason   string
		PostID   int
	}
	var OUTposts []out
	for _, v := range TarReport {
		var temp out
		temp.PostID = v.PostID
		temp.Reason = v.Reason

		var RawUser models.Users
		if err := db.Where("user_id = ?", v.ReportUserID).Find(&RawUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
		temp.Username = RawUser.Username

		var RawPost models.Posts
		if err := db.Where("post_id = ?", v.PostID).Find(&RawPost).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}

		temp.Content = RawPost.Content

		OUTposts = append(OUTposts, temp)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"report_list": OUTposts}, "msg": "success"})
}

func AdminProReport(c *gin.Context) {
	var req models.ApproveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarPost models.Posts
	if err := db.Where("post_id = ?", req.PostID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	TarPost.State = req.Approval
	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 502, "msg": "Internal server error"})
		return
	}

	var TarPost_1 []models.Reports
	if err := db.Where("post_id = ?", req.PostID).Find(&TarPost_1).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 503, "msg": "Internal server error"})
		return
	}

	var newTrash []models.Trash

	// 迁移 TarPost_1 中的数据到 Trash 表
	for _, report := range TarPost_1 {
		trash := models.Trash{
			ReportID:     report.ReportID,
			PostID:       report.PostID,
			ReportUserID: report.ReportUserID,
			Reason:       report.Reason,
			State:        1,          // 设置为1
			CreatedAt:    time.Now(), // 设置为当前时间
		}
		newTrash = append(newTrash, trash)
	}

	if err := db.Create(&newTrash).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 550, "msg": "Internal server error"})
		return
	}

	if err := db.Delete(&TarPost_1).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 504, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}
