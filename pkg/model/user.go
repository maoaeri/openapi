package model

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginInfo struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type User struct {
	UserID    uint   `gorm:"primaryKey" json:"userid"`
	Username  string `json:"username"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
