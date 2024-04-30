package model

import (
	"furnishop/server/enum"
	"time"
)

type User struct {
	Id           string       `gorm:"primaryKey;type:uuid"`
	Name         string       `gorm:"type:varchar(100)"`
	Email        string       `gorm:"unique;type:varchar(100)"`
	Password     string       `gorm:"type:varchar(100)"`
	Role         enum.Role    `gorm:"type:int"`
	Photo        string       `gorm:"type:varchar(100)"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime"`
	ShoppingCart ShoppingCart `gorm:"foreignKey:UserId;references:id"`
}
