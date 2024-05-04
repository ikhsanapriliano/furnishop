package usecasemock

import (
	"furnishop/server/dto"
	"furnishop/server/model"

	"github.com/stretchr/testify/mock"
)

type ProductCategoryUseCaseMock struct {
	mock.Mock
}

func (p *ProductCategoryUseCaseMock) Create(payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	args := p.Called(payload)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryUseCaseMock) GetById(id string) (model.ProductCategory, error) {
	args := p.Called(id)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryUseCaseMock) GetAll() ([]model.ProductCategory, error) {
	args := p.Called()
	return args.Get(0).([]model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryUseCaseMock) Update(id string, payload dto.ProductCategoryDto) (model.ProductCategory, error) {
	args := p.Called(payload)
	return args.Get(0).(model.ProductCategory), args.Error(1)
}

func (p *ProductCategoryUseCaseMock) Delete(id string) error {
	args := p.Called(id)
	return args.Error(0)
}
