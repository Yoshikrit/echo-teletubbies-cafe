package models

import (
    "time"
)

type ProductEntity struct {
    Id 			 int         `gorm:"primaryKey; column:prod_code;"`
    ProdType_Id  int 		 `gorm:"not null;   column:prod_prodtype_code;"`
    Name 		 string      `gorm:"not null;   column:prod_name;              size:40;"`
    Desc 		 string      `gorm:"            column:prod_description;"`
	Price 		 float64	 `gorm:"not null;   column:prod_price;             type:decimal(12,2);"`
	Discount 	 float64	 `gorm:"not null;   column:prod_discount;          type:decimal(5,2);"`
    CreatedAt 	 time.Time   `gorm:"not null;   column:prod_created_at;"`
    CreatedUser  int      	 `gorm:"not null;   column:prod_created_user;"`
    UpdatedAt 	 time.Time   `gorm:"not null;   column:prod_updated_at;"`
    UpdatedUser  int      	 `gorm:"not null;   column:prod_updated_user;"`
}

//make it know table name from database instead of gorm convention
func (p ProductEntity) TableName() string {
	return "product"
}

type Product struct {
    Id     		 int    	`json:"Prod_Id"`
	ProdType_Id  int 	    `json:"ProdType_Id"`
    Name 		 string     `json:"Prod_Name"`
    Desc 		 string     `json:"Prod_Desc"`
	Price 		 float64	`json:"Prod_Price"`
	Discount 	 float64	`json:"Prod_Discount"`
}

type ProductCreate struct {
    Id     		 int    	`json:"Prod_Id"            validate:"required,gt=0"`
	ProdType_Id  int 	    `json:"ProdType_Id"        validate:"required,gt=0"`
    Name 		 string     `json:"Prod_Name"          validate:"required"`
    Desc 		 string     `json:"Prod_Desc"          validate:"omitempty"`
	Price 		 float64	`json:"Prod_Price"         validate:"required,numeric"`
	Discount 	 float64	`json:"Prod_Discount"      validate:"omitempty"`
	CreatedUser  int		`json:"Prod_CreatedUser"   validate:"required,gt=0"`
}

type ProductUpdate struct {
	ProdType_Id  int 	    `json:"ProdType_Id"        validate:"omitempty,gt=0"`
    Name 		 string     `json:"Prod_Name"          validate:"omitempty"`
    Desc 		 string     `json:"Prod_Desc"          validate:"omitempty"`
	Price 		 float64	`json:"Prod_Price"         validate:"omitempty,numeric"`
	Discount 	 float64	`json:"Prod_Discount"      validate:"omitempty"`
	UpdatedUser  int		`json:"Prod_UpdatedUser"   validate:"omitempty,gt=0"`
}
 