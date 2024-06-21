package services_test

import (
	"errors"
	"testing"

	"product/tests/mocks/mock_repositories"
	"product/services"
	"product/utils/errs"
	"github.com/stretchr/testify/assert"

	"product/models"
)

func TestCreateProductType(t *testing.T) {
	prodTypeReqMock := &models.ProductTypeCreate{
		Id:   1,
		Name: "A",
	}

	prodTypeResMock := &models.ProductType{
		Id:   1,
		Name: "A",
	}

	prodTypeFromDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("Create", prodTypeReqMock).Return(prodTypeFromDBMock, nil)

		service := services.NewProductTypeService(mockRepo)
		prodTypeRes, err := service.CreateProductType(prodTypeReqMock)

		expected := prodTypeResMock
		assert.NoError(t, err)
		assert.Equal(t, expected, prodTypeRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	type testCase struct {
		test_name       string
		isNull          bool
		id              int	 	
		name            string
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no id",   	isNull: true, 	id: 0, name: "",    err: errs.NewBadRequestError("ProductType's Id is null")},
		{test_name: "test case : fail no name", 	isNull: true, 	id: 1, name:  "",   err: errs.NewBadRequestError("ProductType's Name is null")},
		{test_name: "test case : fail repository", 	isNull: false, id: 1, name:  "A",  err: errors.New("")},
	}

	for _, tc := range cases {
		prodTypeReqFail := &models.ProductTypeCreate{
			Id:   tc.id,
			Name: tc.name,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			if !tc.isNull {
				mockRepo.On("Create", prodTypeReqFail).Return(&models.ProductTypeEntity{}, errors.New(""))
			}
			service := services.NewProductTypeService(mockRepo)

			prodTypeRes, err := service.CreateProductType(prodTypeReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, prodTypeRes)
			if !tc.isNull {
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestGetProductTypes(t *testing.T) {
	prodTypesDBMock := []models.ProductTypeEntity{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	prodTypesResMock := []models.ProductType{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetAll").Return(prodTypesDBMock, nil)

		service := services.NewProductTypeService(mockRepo)
		prodTypesRes, err := service.GetProductTypes()

		assert.NoError(t, err)
		assert.Equal(t, prodTypesResMock, prodTypesRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetAll").Return([]models.ProductTypeEntity{}, errors.New(""))

		service := services.NewProductTypeService(mockRepo)
		prodTypesRes, err := service.GetProductTypes()

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodTypesRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProductType(t *testing.T) {
	prodTypeDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "A",
	}
	prodTypeResMock := &models.ProductType{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetById", 1).Return(prodTypeDBMock, nil)

		service := services.NewProductTypeService(mockRepo)
		prodTypeResponse, err := service.GetProductType(1)

		assert.NoError(t, err)
		assert.Equal(t, prodTypeResMock, prodTypeResponse)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetById", 1).Return(&models.ProductTypeEntity{}, errors.New(""))

		service := services.NewProductTypeService(mockRepo)
		prodTypeRes, err := service.GetProductType(1)

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodTypeRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProductType(t *testing.T) {
	prodTypeReqMock := &models.ProductTypeUpdate{
		Name: "B",
	}
	prodTypeDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "B",
	}
	prodTypeResMock := &models.ProductType{
		Id:   1,
		Name: "B",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("Update", 1, prodTypeReqMock).Return(prodTypeDBMock, nil)

		service := services.NewProductTypeService(mockRepo)
		prodTypeRes, err := service.UpdateProductType(1, prodTypeReqMock)

		assert.NoError(t, err)
		assert.Equal(t, prodTypeResMock, prodTypeRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	type testCase struct {
		test_name       string
		isNull 			bool	
		name            string
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no name",  	isNull: true, 	name:  "",   err: errs.NewBadRequestError("ProductType's Name is null")},
		{test_name: "test case : fail repository",  isNull: false, 	name:  "B",  err: errors.New("")},
	}

	for _, tc := range cases {
		prodTypeReqFail := &models.ProductTypeUpdate{
			Name: tc.name,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			if !tc.isNull {
				mockRepo.On("Update", 1, prodTypeReqMock).Return(&models.ProductTypeEntity{}, errors.New(""))
			}
			service := services.NewProductTypeService(mockRepo)

			prodTypeRes, err := service.UpdateProductType(1, prodTypeReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, prodTypeRes)
			if !tc.isNull {
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestDeleteProductType(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(nil)

		service := services.NewProductTypeService(mockRepo)
		err := service.DeleteProductType(1)

		assert.NoError(t, err)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(errors.New(""))

		service := services.NewProductTypeService(mockRepo)
		err := service.DeleteProductType(1)

		expected := errors.New("")
		assert.Equal(t, expected, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProductTypeCount(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetCount").Return(int64(5), nil)

		service := services.NewProductTypeService(mockRepo)
		count, err := service.GetProductTypeCount()

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, int64(5), count)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetCount").Return(int64(0), errors.New(""))

		service := services.NewProductTypeService(mockRepo)
		count, err := service.GetProductTypeCount()

		expected := errors.New("")
		assert.Equal(t, expected, err)
		assert.Equal(t, int64(0), count)
		mockRepo.AssertExpectations(t)
	})
}
