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

func TestCreateProduct(t *testing.T) {
	prodReqMock := &models.ProductCreate{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Desc: "A",
		Price: 100,
		Discount: 0,
		CreatedUser: 1,
	}

	prodResMock := &models.Product{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}

	prodFromDBMock := &models.ProductEntity{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("Create", prodReqMock).Return(prodFromDBMock, nil)

		service := services.NewProductService(mockRepo)
		prodRes, err := service.CreateProduct(prodReqMock)

		expected := prodResMock
		assert.NoError(t, err)
		assert.Equal(t, expected, prodRes)
		assert.Nil(t, err)
	})

	type testCase struct {
		test_name       string
		isNull     		 bool	 	
		id              int	 	
		typeId          int	 
		name            string
		price           float64
		createdUser     int	 
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no id",       	isNull: true, 	id: 0, typeId: 0, name:  "",   price: 0,  createdUser: 0,   err: errs.NewBadRequestError("Product's Id is null")},
		{test_name: "test case : fail no prod id",  	isNull: true, 	id: 1, typeId: 0, name:  "",   price: 0,  createdUser: 0,   err: errs.NewBadRequestError("ProductType's Id is null")},
		{test_name: "test case : fail no name",     	isNull: true, 	id: 1, typeId: 1, name:  "",   price: 0,  createdUser: 0,   err: errs.NewBadRequestError("Product's Name is null")},
		{test_name: "test case : fail no price", 		isNull: true, 	id: 1, typeId: 1, name:  "A",  price: 0,  createdUser: 0,   err: errs.NewBadRequestError("Product's Price is null")},
		{test_name: "test case : fail no createduser", 	isNull: true, 	id: 1, typeId: 1, name:  "A",  price: 1,  createdUser: 0,   err: errs.NewBadRequestError("Product's CreatedUser is null")},
		{test_name: "test case : fail repository",      isNull: false, 	id: 1, typeId: 1, name:  "A",  price: 1,  createdUser: 1,   err: errors.New("")},
	}

	for _, tc := range cases {
		prodReqErrorMock := &models.ProductCreate{
			Id:          tc.id,
			ProdType_Id: tc.typeId,
			Name:        tc.name,
			Price:       tc.price,
			Discount:    0,
			CreatedUser: tc.createdUser,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewProductRepositoryMock()
			if !tc.isNull {
				mockRepo.On("Create", prodReqErrorMock).Return(&models.ProductEntity{}, errors.New(""))
			}

			service := services.NewProductService(mockRepo)
			prodRes, err := service.CreateProduct(prodReqErrorMock)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, prodRes)
			if !tc.isNull {
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestGetProducts(t *testing.T) {
	prodsDBMock := []models.ProductEntity{
		{
			Id:   1,
			ProdType_Id:   1,
			Name: "A",
			Desc: "A",
			Price: 100,
			Discount: 0,
		},
		{
			Id:   2,
			ProdType_Id:   1,
			Name: "B",
			Desc: "B",
			Price: 100,
			Discount: 0,
		},
	}

	prodsResMock := []models.Product{
		{
			Id:   1,
			ProdType_Id:   1,
			Name: "A",
			Desc: "A",
			Price: 100,
			Discount: 0,
		},
		{
			Id:   2,
			ProdType_Id:   1,
			Name: "B",
			Desc: "B",
			Price: 100,
			Discount: 0,
		},
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetAll").Return(prodsDBMock, nil)

		service := services.NewProductService(mockRepo)
		prodsRes, err := service.GetProducts()

		assert.NoError(t, err)
		assert.Equal(t, prodsResMock, prodsRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetAll").Return([]models.ProductEntity{}, errors.New(""))

		service := services.NewProductService(mockRepo)
		prodsRes, err := service.GetProducts()

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodsRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProduct(t *testing.T) {
	prodDBMock := &models.ProductEntity{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Price: 100,
		Discount: 0,
	}
	prodResMock := &models.Product{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Price: 100,
		Discount: 0,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetById", 1).Return(prodDBMock, nil)

		service := services.NewProductService(mockRepo)
		prodResponse, err := service.GetProduct(1)

		assert.NoError(t, err)
		assert.Equal(t, prodResMock, prodResponse)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetById", 1).Return(&models.ProductEntity{}, errors.New(""))

		service := services.NewProductService(mockRepo)
		prodRes, err := service.GetProduct(1)

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	prodReqMock := &models.ProductUpdate{
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
		UpdatedUser: 1,
	}
	prodReqErrorMock := &models.ProductUpdate{
		ProdType_Id:   0,
		Name: "",
		Desc: "",
		Price: 0,
		Discount: 0,
		UpdatedUser: 0,
	}
	prodDBMock := &models.ProductEntity{
		Id:   1,
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}
	prodResMock := &models.Product{
		Id:   1,
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("Update", 1, prodReqMock).Return(prodDBMock, nil)

		service := services.NewProductService(mockRepo)
		prodRes, err := service.UpdateProduct(1, prodReqMock)

		assert.NoError(t, err)
		assert.Equal(t, prodResMock, prodRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail null", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()

		service := services.NewProductService(mockRepo)
		prodRes, err := service.UpdateProduct(1, prodReqErrorMock)

		expected := errs.NewBadRequestError("Product's Data is null")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodRes)
	})
	t.Run("test case : fail repository", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("Update", 1, prodReqMock).Return(prodDBMock, errors.New(""))

		service := services.NewProductService(mockRepo)
		prodRes, err := service.UpdateProduct(1, prodReqMock)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(nil)

		service := services.NewProductService(mockRepo)
		err := service.DeleteProduct(1)

		assert.NoError(t, err)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(errors.New(""))

		service := services.NewProductService(mockRepo)
		err := service.DeleteProduct(1)

		expected := errors.New("")
		assert.Equal(t, expected, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProductCount(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetCount").Return(int64(5), nil)

		service := services.NewProductService(mockRepo)
		count, err := service.GetProductCount()

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, int64(5), count)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewProductRepositoryMock()
		mockRepo.On("GetCount").Return(int64(0), errors.New(""))

		service := services.NewProductService(mockRepo)
		count, err := service.GetProductCount()

		expected := errors.New("")
		assert.Equal(t, expected, err)
		assert.Equal(t, int64(0), count)
		mockRepo.AssertExpectations(t)
	})
}