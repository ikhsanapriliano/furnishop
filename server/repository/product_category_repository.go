package repository

import (
	"furnishop/server/dto"
	"furnishop/server/model"

	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	Create(payload model.ProductCategory) (model.ProductCategory, error)
	GetById(id string) (model.ProductCategory, error)
	GetAll() ([]model.ProductCategory, error)
	Update(data model.ProductCategory, payload dto.ProductCategoryDto) (model.ProductCategory, error)
	Delete(payload model.ProductCategory) error
}

type productCategoryRepository struct {
	db *gorm.DB
}

func (p *productCategoryRepository) Create(payload model.ProductCategory) (model.ProductCategory, error) {
	err := p.db.Create(&payload).Error
	if err != nil {
		return model.ProductCategory{}, err
	}

	return payload, nil
}

func (p *productCategoryRepository) GetById(id string) (model.ProductCategory, error) {
	var data model.ProductCategory
	err := p.db.First(&data, id).Error
	if err != nil {
		return model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryRepository) GetAll() ([]model.ProductCategory, error) {
	var data []model.ProductCategory
	err := p.db.Find(&data).Error
	if err != nil {
		return []model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryRepository) Update(data model.ProductCategory, payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	err := p.db.Model(&data).Updates(payload).Error
	if err != nil {
		return model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryRepository) Delete(payload model.ProductCategory) error {
	err := p.db.Delete(&payload).Error
	if err != nil {
		return err
	}

	return nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		db: db,
	}
}
