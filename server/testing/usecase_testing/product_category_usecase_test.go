package usecasetesting

import (
	"furnishop/server/dto"
	repomock "furnishop/server/mock/repo_mock"
	"furnishop/server/model"
	"furnishop/server/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductCategoryUseCaseTestSuite struct {
	suite.Suite
	pcrm *repomock.ProductCategoryRepoMock
	pcu  usecase.ProductCategoryUseCase
}

func (suite *ProductCategoryUseCaseTestSuite) SetupTest() {
	suite.pcrm = new(repomock.ProductCategoryRepoMock)
	suite.pcu = usecase.NewProductCategoryUseCase(suite.pcrm)
}

func TestProductCategoryUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProductCategoryUseCaseTestSuite))
}

func (suite *ProductCategoryUseCaseTestSuite) TestCreate_Success() {
	payloadMock := dto.ProductCategoryDto{
		Name: "1",
	}

	expected := model.ProductCategory{
		Name: payloadMock.Name,
	}

	suite.pcrm.Mock.On("Create", mock.Anything).Return(expected, nil)

	actual, err := suite.pcu.Create(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Name, actual.Name)
}

func (suite *ProductCategoryUseCaseTestSuite) TestGetById_Success() {
	idMock := "1"
	expected := model.ProductCategory{
		Id: idMock,
	}

	suite.pcrm.Mock.On("GetById", mock.Anything).Return(expected, nil)

	actual, err := suite.pcu.GetById(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *ProductCategoryUseCaseTestSuite) TestGetAll_Success() {
	expected := []model.ProductCategory{
		{
			Id: "1",
		},
	}

	suite.pcrm.Mock.On("GetAll").Return(expected, nil)

	actual, err := suite.pcu.GetAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected[0].Id, actual[0].Id)
}

func (suite *ProductCategoryUseCaseTestSuite) TestUpdate_Success() {
	idMock := "1"
	payloadMock := dto.ProductCategoryDto{
		Name: "1",
	}

	expected := model.ProductCategory{
		Id:   idMock,
		Name: payloadMock.Name,
	}

	suite.pcrm.Mock.On("GetById", mock.Anything).Return(model.ProductCategory{}, nil)
	suite.pcrm.Mock.On("Update", mock.Anything, mock.Anything).Return(expected, nil)

	actual, err := suite.pcu.Update(idMock, payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
	assert.Equal(suite.T(), expected.Name, actual.Name)
}

func (suite *ProductCategoryUseCaseTestSuite) TestDelete_Success() {
	idMock := "1"

	suite.pcrm.Mock.On("GetById", mock.Anything).Return(model.ProductCategory{}, nil)
	suite.pcrm.Mock.On("Delete", mock.Anything, mock.Anything).Return(nil)

	err := suite.pcu.Delete(idMock)
	assert.Nil(suite.T(), err)
}
