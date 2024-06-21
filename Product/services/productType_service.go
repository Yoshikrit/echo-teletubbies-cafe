package services

import (
	"product/models"
	"product/repositories"

	"product/utils/errs"
	"product/utils/logs"
)

type productTypeService struct {
	prodTypeRepo repositories.ProductTypeRepository
}

func NewProductTypeService(prodTypeRepo repositories.ProductTypeRepository) ProductTypeService {
	return productTypeService{prodTypeRepo: prodTypeRepo}
}

func (s productTypeService) CreateProductType(prodTypeReq *models.ProductTypeCreate) (*models.ProductType, error) {
	if prodTypeReq.Id == 0{
		return nil, errs.NewBadRequestError("ProductType's Id is null")
	}
	if prodTypeReq.Name == ""{
		return nil, errs.NewBadRequestError("ProductType's Name is null")
	}

	prodTypeEntityRes, err := s.prodTypeRepo.Create(prodTypeReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	
	prodTypeRes := models.ProductType{
		Id:       prodTypeEntityRes.Id,
		Name:     prodTypeEntityRes.Name,
	}

	logs.Info("Service: Create ProductType Successfully")
	return &prodTypeRes, nil
}

func (s productTypeService) GetProductTypes() ([]models.ProductType, error) {
	prodTypesFromDB, err := s.prodTypeRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var prodTypesRes []models.ProductType
	for _, prodTypeFromDB := range prodTypesFromDB {
		prodTypeRes := models.ProductType{
			Id:       prodTypeFromDB.Id,
			Name:     prodTypeFromDB.Name,
		}
		prodTypesRes = append(prodTypesRes, prodTypeRes)
	}

	logs.Info("Service: Get ProductTypes Successfully")
	return prodTypesRes, nil
}

func (s productTypeService) GetProductType(id int) (*models.ProductType, error) {
	prodTypeFromDB, err := s.prodTypeRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	prodTypeRes := models.ProductType{
		Id:       prodTypeFromDB.Id,
		Name:     prodTypeFromDB.Name,
	}

	logs.Info("Service: Get ProductType Successfully")
	return &prodTypeRes, nil
}

func (s productTypeService) UpdateProductType(id int, reqProdType *models.ProductTypeUpdate) (*models.ProductType, error) {
	if reqProdType.Name == ""{
		return nil, errs.NewBadRequestError("ProductType's Name is null")
	}

	prodTypeFromDB, err := s.prodTypeRepo.Update(id, reqProdType)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	prodTypeRes := models.ProductType{
		Id:       prodTypeFromDB.Id,
		Name:     prodTypeFromDB.Name,
	}

	logs.Info("Service: Update ProductType Successfully")
	return &prodTypeRes, nil
}

func (s productTypeService) DeleteProductType(id int) (error) {
	err := s.prodTypeRepo.DeleteById(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	
	logs.Info("Service: Delete ProductType Successfully")
	return nil
}

func (s productTypeService) GetProductTypeCount() (int64, error) {
	count, err := s.prodTypeRepo.GetCount()
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get ProductType's Count Successfully")
	return count, nil
}

