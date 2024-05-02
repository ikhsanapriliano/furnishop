package factory

import (
	"furnishop/server/model"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCatgeoryFactory struct {
	db    *gorm.DB
	count int
}

func (p *ProductCatgeoryFactory) Seed() {
	for i := 0; i < p.count; i++ {
		data := model.ProductCategory{
			Id:   uuid.NewString(),
			Name: faker.LastName(),
		}

		err := p.db.Save(&data).Error
		if err != nil {
			panic(err)
		}
	}
}

func NewProductCategoryFactory(db *gorm.DB, count int) *ProductCatgeoryFactory {
	return &ProductCatgeoryFactory{
		db:    db,
		count: count,
	}
}
