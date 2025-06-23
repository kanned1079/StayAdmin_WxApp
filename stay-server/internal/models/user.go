package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 用户id
	PlatformId string         `json:"platform_id"`
	AppOpenId  string         `json:"app_open_id" gorm:"unique"`
	SessionKey string         `json:"session_key"`
	Name       string         `json:"name"`
	Role       string         `json:"role"`
	CreatedAt  time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  time.Time      `json:"Updated_at" gorm:"type:timestamp"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"type:timestamp"`
}

func (User) TableName() string {
	return "s_user"
}
