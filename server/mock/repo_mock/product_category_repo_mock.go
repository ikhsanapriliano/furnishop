package repomock

import (
	"furnishop/server/dto"
	"furnishop/server/model"

	"github.com/stretchr/testify/mock"
)

type ProductCategoryRepoMock struct {
	mock.Mock
}

func (p *ProductCategoryRepoMock) Create(payload model.ProductCategory) (model.ProductCategory, error) {
	args := p.Called(payload)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryRepoMock) GetById(id string) (model.ProductCategory, error) {
	args := p.Called(id)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryRepoMock) GetAll() ([]model.ProductCategory, error) {
	args := p.Called()
	return args.Get(0).([]model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryRepoMock) Update(data model.ProductCategory, payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	args := p.Called(data, payload)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryRepoMock) Delete(payload model.ProductCategory) error {
	args := p.Called(payload)
	return args.Error(0)
}
