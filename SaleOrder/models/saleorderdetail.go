package models

import (

)

type SaleOrderDetailEntity struct {
	Seq       int       `gorm:"autoIncrement; primaryKey; column:sod_seq"`
	SO_Id     int       `gorm:"not null; column:sod_so_code"`
	Prod_Id   int       `gorm:"not null; column:sod_prod_code"`
	Quantity  int       `gorm:"not null; column:sod_quantity"`
	Price     float64   `gorm:"not null; column:sod_price; type:decimal(12,2)"`
	Discount  float64   `gorm:"not null; column:sod_discount; type:decimal(5,2)"`
}

// TableName sets the table name for the entity
func (s *SaleOrderDetailEntity) TableName() string {
	return "saleorderdetail"
}


type SaleOrderDetail struct {
    Seq     	 int    	`json:"SOD_Id"`
	SO_Id  		 int 	    `json:"SO_Id"`
    Prod_Id  	 int  		`json:"Prod_Id"`
    Quantity 	 int    	`json:"SOD_Qty"`
	Price    	 float64	`json:"SOD_Price"`
	Discount 	 float64	`json:"SOD_Discount"`
}

type SaleOrderDetailCreate struct {
	SO_Id  		 int 	    `json:"SO_Id"              validate:"required,gt=0"`
    Prod_Id  	 int  		`json:"Prod_Id"            validate:"required,gt=0"`
    Quantity 	 int    	`json:"SOD_Qty"            validate:"required,gt=0"`
	Price    	 float64	`json:"SOD_Price"          validate:"required,numeric"`
	Discount 	 float64	`json:"SOD_Discount"       validate:"omitempty,numeric"`
}

type SaleOrderDetailRate struct {
	Prod_Id  	 int 	    `json:"Prod_Id"       gorm:"column:sod_prod_code"`
    Quantity 	 int    	`json:"SOD_Qty"`
	Price    	 float64	`json:"SOD_Price"`
}
