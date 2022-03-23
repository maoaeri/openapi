package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	PostID    int    `gorm:"primaryKey" json:"postid"`
	UserID    int    `json:"userid"`
	Content   string `json:"content"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
