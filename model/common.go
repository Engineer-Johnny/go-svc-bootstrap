package model

import (
	"time"

	"gorm.io/gorm"
)

type COMMON_MODEL struct {
	ID        uint   `gorm:"primarykey"`
	UUID      string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:UUID;type:varchar(100);size:100;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
