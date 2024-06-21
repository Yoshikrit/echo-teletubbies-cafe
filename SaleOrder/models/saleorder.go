package models

import (
	"time"
)

type SaleOrderEntity struct {
    Id 			 int        `gorm:"autoIncrement; primaryKey; column:saleorder_code;"`
    CreatedUser  int        `gorm:"not null;      column:saleorder_created_user;"`
	CreatedDate  time.Time  `gorm:"not null;      column:saleorder_created_at;"`
    TotalPrice   float64    `gorm:"               column:saleorder_total_price;            type:decimal(12,2);"`
    PayMethod    int        `gorm:"               column:saleorder_paymentchannel_code;"`
	Status       string     `gorm:"not null;      column:saleorder_status;                 size:4;"`
}

//make it know table name from database instead of gorm convention
func (p SaleOrderEntity) TableName() string {
	return "saleorder"
}

type SaleOrder struct {
    Id     		 int    	`json:"SO_Id"`
	CreatedUser  int 	    `json:"User_Id"`
    CreatedDate  time.Time  `json:"SO_CreatedDate"`
    TotalPrice 	 float64    `json:"SO_TotalPrice"`
	Status 	     string	    `json:"SO_Status"`
	PayMethod    int	    `json:"SO_PayMethod"`
}

type SaleOrderCreate struct {
	CreatedUser  int 	    `json:"User_Id"             validate:"required,gt=0"`
    TotalPrice 	 float64    `json:"SO_TotalPrice"       validate:"omitempty,numeric"`
	Status       string     `json:"SO_Status"           validate:"required"`
	PayMethod    int        `json:"SO_PayMethod"        validate:"omitempty"`
}

type SaleOrderReport struct {
	Seq            int        `json:"SO_Seq"`  
    Id     		   int    	  `json:"SO_Id"`
	User           string 	  `json:"SO_EMP_Name"`
    Date           time.Time  `json:"SO_Date"`
    TotalPrice 	   float64    `json:"SO_TotalPrice"`
	Status 	       string	  `json:"SO_Status"`
	PayMethodName  string	  `json:"PayMethod_Name"`
}