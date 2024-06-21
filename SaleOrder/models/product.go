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