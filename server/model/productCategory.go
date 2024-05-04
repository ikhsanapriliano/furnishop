package model

import (
	"time"
)

type ProductCategory struct {
	Id        string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `gorm:"unique;type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
