package models

import (
	"gorm.io/gorm"
	"time"
)

type Station struct {
	Id          int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 数据表id
	StationName string         `json:"station_name"`
	Address     string         `json:"address"`
	MasterId    int64          `json:"master_id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time      `json:"Updated_at" gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"type:timestamp"`
}

func (Station) TableName() string {
	return "s_station"
}
