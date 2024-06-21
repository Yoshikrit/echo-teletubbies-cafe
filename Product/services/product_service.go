package services

import (
	"product/models"
	"product/repositories"

	"product/utils/errs"
	"product/utils/logs"
)

type productService struct {
	prodRepo repositories.ProductRepository
}

func NewProductService(prodRepo repositories.ProductRepository) ProductService {
	return productService{prodRepo: prodRepo}
}

func (s productService) CreateProduct(prodReq *models.ProductCreate) (*models.Product, error) {
	if prodReq.Id == 0{
		return nil, errs.NewBadRequestError("Product's Id is null")
	}
	if prodReq.ProdType_Id == 0{
		return nil, errs.NewBadRequestError("ProductType's Id is null")
	}
	if prodReq.Name == ""{
		return nil, errs.NewBadRequestError("Product's Name is null")
	}
	if prodReq.Price == 0{
		return nil, errs.NewBadRequestError("Product's Price is null")
	}
	if prodReq.CreatedUser == 0{
		return nil, errs.NewBadRequestError("Product's CreatedUser is null")
	}

	prodEntityRes, err := s.prodRepo.Create(prodReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	
	prodRes := models.Product{
		Id:       		prodEntityRes.Id,
		ProdType_Id:    prodEntityRes.ProdType_Id,
		Name:     		prodEntityRes.Name,
		Desc:     		prodEntityRes.Desc,
		Price:     		prodEntityRes.Price,
		Discount:     	prodEntityRes.Discount,
	}

	logs.Info("Service: Create Product Successfully")
	return &prodRes, nil
}

func (s productService) GetProducts() ([]models.Product, error) {
	prodsFromDB, err := s.prodRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var prodsRes []models.Product
	for _, prodFromDB := range prodsFromDB {
		prodRes := models.Product{
			Id:       		prodFromDB.Id,
			ProdType_Id:    prodFromDB.ProdType_Id,
			Name:     		prodFromDB.Name,
			Desc:			prodFromDB.Desc,
			Price:     		prodFromDB.Price,
			Discount:     	prodFromDB.Discount,
		}
		prodsRes = append(prodsRes, prodRes)
	}

	logs.Info("Service: Get Products Successfully")
	return prodsRes, nil
}

func (s productService) GetProduct(id int) (*models.Product, error) {
	prodFromDB, err := s.prodRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	prodRes := models.Product{
		Id:       		prodFromDB.Id,
		ProdType_Id:    prodFromDB.ProdType_Id,
		Name:     		prodFromDB.Name,
		Desc:     		prodFromDB.Desc,
		Price:     		prodFromDB.Price,
		Discount:     	prodFromDB.Discount,
	}

	logs.Info("Service: Get Product Successfully")
	return &prodRes, nil
}

func (s productService) UpdateProduct(id int, prodReq *models.ProductUpdate) (*models.Product, error) {
	if (prodReq.ProdType_Id == 0 && prodReq.Name == "" && prodReq.Price == 0 && prodReq.UpdatedUser == 0 && prodReq.Desc == "" && prodReq.Discount == 0) {
		return nil, errs.NewBadRequestError("Product's Data is null")
	}

	prodFromDB, err := s.prodRepo.Update(id, prodReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	prodRes := models.Product{
		Id:       		prodFromDB.Id,
		ProdType_Id:    prodFromDB.ProdType_Id,
		Name:     		prodFromDB.Name,
		Desc:     		prodFromDB.Desc,
		Price:     		prodFromDB.Price,
		Discount:     	prodFromDB.Discount,
	}

	logs.Info("Service: Update Product Successfully")
	return &prodRes, nil
}

func (s productService) DeleteProduct(id int) (error) {
	err := s.prodRepo.DeleteById(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	
	logs.Info("Service: Delete Product Successfully")
	return nil
}

func (s productService) GetProductCount() (int64, error) {
	count, err := s.prodRepo.GetCount()
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Product's Count Successfully")
	return count, nil
}

