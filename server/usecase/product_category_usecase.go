package usecase

import (
	"furnishop/server/dto"
	"furnishop/server/model"
	"furnishop/server/repository"

	"github.com/google/uuid"
)

type ProductCategoryUseCase interface {
	Create(payload dto.ProductCategoryDto) (model.ProductCategory, error)
	GetById(id string) (model.ProductCategory, error)
	GetAll() ([]model.ProductCategory, error)
	Update(id string, payload dto.ProductCategoryDto) (model.ProductCategory, error)
	Delete(id string) error
}

type productCategoryUseCase struct {
	repo repository.ProductCategoryRepository
}

func (p *productCategoryUseCase) Create(payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	input := model.ProductCategory{
		Id:   uuid.NewString(),
		Name: payload.Name,
	}
	data, err := p.repo.Create(input)
	if err != nil {
		return model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryUseCase) GetById(id string) (model.ProductCategory, error) {
	data, err := p.repo.GetById(id)
	if err != nil {
		return model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryUseCase) GetAll() ([]model.ProductCategory, error) {
	data, err := p.repo.GetAll()
	if err != nil {
		return []model.ProductCategory{}, err
	}

	return data, nil
}

func (p *productCategoryUseCase) Update(id string, payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	data, err := p.repo.GetById(id)
	if err != nil {
		return model.ProductCategory{}, err
	}

	newData, err := p.repo.Update(data, payload)
	if err != nil {
		return model.ProductCategory{}, err
	}

	return newData, nil
}

func (p *productCategoryUseCase) Delete(id string) error {
	data, err := p.repo.GetById(id)
	if err != nil {
		return err
	}

	err = p.repo.Delete(data)
	if err != nil {
		return err
	}

	return nil
}

func NewProductCategoryUseCase(repo repository.ProductCategoryRepository) ProductCategoryUseCase {
	return &productCategoryUseCase{
		repo: repo,
	}
}
