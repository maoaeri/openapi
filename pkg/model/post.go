package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	PostID    uint   `gorm:"primaryKey" json:"postid"`
	UserID    uint   `json:"userid"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
