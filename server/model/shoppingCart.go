package model

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type ShoppingCart struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	UserId    uuid.UUID `gorm:"unique;type:uuid" json:"userId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
