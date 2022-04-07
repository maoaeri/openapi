package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	CommentID int    `gorm:"primaryKey" json:"commentid"`
	UserID    int    `json:"userid"`
	Content   string `json:"content"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
