package repotesting

import (
	"furnishop/server/dto"
	"furnishop/server/model"
	"furnishop/server/repository"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductCategoryRepoTestSuite struct {
	suite.Suite
	mockDb  *gorm.DB
	mockSql sqlmock.Sqlmock
	repo    repository.ProductCategoryRepository
}

func (suite *ProductCategoryRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	suite.mockDb, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.Nil(suite.T(), err)
	suite.mockSql = mock
	suite.repo = repository.NewProductCategoryRepository(suite.mockDb)
}

func TestProductCategoryRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductCategoryRepoTestSuite))
}

func newTable() *sqlmock.Rows {
	var table = sqlmock.NewRows([]string{
		"Id",
		"Name",
		"CreatedAt",
		"UpdatedAt",
	})

	return table
}

func (suite *ProductCategoryRepoTestSuite) TestCreate_Success() {
	payloadMock := model.ProductCategory{
		Id: "1",
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec(regexp.QuoteMeta(`INSERT`)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Create(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), payloadMock.Id, actual.Id)
}

func (suite *ProductCategoryRepoTestSuite) TestGetById_Success() {
	idMock := "1"
	expected := model.ProductCategory{
		Id:        idMock,
		Name:      "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	table := newTable()
	table.AddRow(expected.Id, expected.Name, expected.CreatedAt, expected.UpdatedAt)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(table)

	actual, err := suite.repo.GetById(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *ProductCategoryRepoTestSuite) TestGetAll_Success() {
	expected := []model.ProductCategory{
		{
			Id:        "1",
			Name:      "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	table := newTable()
	table.AddRow(expected[0].Id, expected[0].Name, expected[0].CreatedAt, expected[0].UpdatedAt)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(table)

	actual, err := suite.repo.GetAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected[0].Id, actual[0].Id)
}

func (suite *ProductCategoryRepoTestSuite) TestUpdate_Success() {
	dataMock := model.ProductCategory{
		Id:   "1",
		Name: "1",
	}
	payloadMock := dto.ProductCategoryDto{
		Name: "2",
	}

	expected := model.ProductCategory{
		Id:   dataMock.Id,
		Name: payloadMock.Name,
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Update(dataMock, payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Name, actual.Name)
}

func (suite *ProductCategoryRepoTestSuite) TestDelete_Success() {
	dataMock := model.ProductCategory{
		Id: "1",
	}

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec(regexp.QuoteMeta(`DELETE`)).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()

	err := suite.repo.Delete(dataMock)
	assert.Nil(suite.T(), err)
}
