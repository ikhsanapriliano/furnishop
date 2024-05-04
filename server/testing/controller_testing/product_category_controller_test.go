package controllertesting

import (
	"bytes"
	"encoding/json"
	"furnishop/server/delivery/controller"
	"furnishop/server/dto"
	usecasemock "furnishop/server/mock/usecase_mock"
	"furnishop/server/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductCategoryControllerTestSuite struct {
	suite.Suite
	pcum       *usecasemock.ProductCategoryUseCaseMock
	rg         *gin.RouterGroup
	controller *controller.ProductCategoryController
}

func (suite *ProductCategoryControllerTestSuite) SetupTest() {
	suite.pcum = new(usecasemock.ProductCategoryUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.controller = controller.NewProductCategoryController(suite.pcum, suite.rg)
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductCategoryControllerTestSuite))
}

func (suite *ProductCategoryControllerTestSuite) TestCreateHandler_Success() {
	suite.pcum.Mock.On("Create", mock.Anything).Return(model.ProductCategory{}, nil)
	payloadMock := dto.ProductCategoryDto{
		Name: "1",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/prodctCategories", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.CreateHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *ProductCategoryControllerTestSuite) TestGetByIdHandler_Success() {
	suite.pcum.Mock.On("GetById", mock.Anything).Return(model.ProductCategory{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/prodctCategories/:id", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetByIdHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *ProductCategoryControllerTestSuite) TestGetAll_Success() {
	suite.pcum.Mock.On("GetAll", mock.Anything).Return([]model.ProductCategory{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/prodctCategories", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetAllHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *ProductCategoryControllerTestSuite) TestUpdate_Success() {
	suite.pcum.Mock.On("Update", mock.Anything).Return(model.ProductCategory{}, nil)
	payloadMock := dto.ProductCategoryDto{
		Name: "1",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/prodctCategories/{id}", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.UpdateHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *ProductCategoryControllerTestSuite) TestDelete_Success() {
	suite.pcum.Mock.On("Delete", mock.Anything).Return(nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/prodctCategories/{id}", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.DeleteHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}
