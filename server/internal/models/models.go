package models

import (
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}

type PostRequest struct {
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

type RePostRequest struct {
	Content string `json:"content"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
}

type DePostRequest struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}

type ReportPostRequest struct {
	PostID int    `json:"post_id"`
	Reason string `json:"reason"`
	UserID int    `json:"user_id"`
}

type ApproveRequest struct {
	Approval int `json:"approval"`
	PostID   int `json:"post_id"`
	UserID   int `json:"user_id"`
}

type Users struct {
	UserID    int       `gorm:"column:user_id;primaryKey;autoIncrement"`    // 对应 user_id 列，设置为主键和自增
	Username  string    `gorm:"column:username;unique;size:50"`             // 对应 username 列，唯一，长度为 50
	Name      string    `gorm:"column:name;size:100"`                       // 对应 name 列，长度为 100
	Password  string    `gorm:"column:password;size:255"`                   // 对应 password 列，长度为 255
	UserType  int       `gorm:"column:user_type;check:user_type IN (1, 2)"` // 对应 user_type 列，并限制其值为 1 或 2
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`           // 对应 created_at 列，自动设置创建时间
}

type Posts struct {
	PostID    int       `gorm:"column:post_id;primaryKey;autoIncrement"`     // 对应 post_id 列，设置为主键和自增
	UserID    int       `gorm:"column:user_id"`                              // 对应 user_id 列
	Username  string    `gorm:"column:username;size:50"`                     // 对应 username 列，唯一，长度为 50
	Name      string    `gorm:"column:name;size:100"`                        // 对应 name 列，长度为 100
	UserType  int       `gorm:"column:user_type;check:user_type IN (1, 2)"`  // 对应 user_type 列，并限制其值为 1 或 2
	Content   string    `gorm:"column:content;size:255"`                     // 对应 content 列，长度为 255
	Reason    string    `gorm:"column:reason;size:255"`                      // 对应 reason 列，长度为 255
	State     int       `gorm:"column:state"`                                // 对应 state 列
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp"` // 对应 created_at 列，默认值为当前时间戳
}

type Reports struct {
	ReportID     int       `gorm:"column:report_id;primaryKey;autoIncrement"`
	PostID       int       `gorm:"column:post_id"`
	ReportUserID int       `gorm:"column:report_user_id"`
	Reason       string    `gorm:"column:reason"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type Trash struct {
	ReportID int       `gorm:"column:report_id;type:int;primaryKey"`
	PostID   int       `gorm:"column:post_id;type:int"`
	ReportUserID int    `gorm:"column:report_user_id;type:int"`
	Reason    string   `gorm:"column:reason;type:varchar(255)"`
	State     int      `gorm:"column:state;type:int"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
}