package database

import (
	"furnishop/server/config"
	"furnishop/server/manager"
	"log"

	"gorm.io/gorm"
)

type Initializer struct {
}

func (i *Initializer) Start() *gorm.DB {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, _ := manager.NewInfraManager(cfg)
	db := infraManager.Conn()

	return db
}

func NewInitializer() *Initializer {
	return &Initializer{}
}
