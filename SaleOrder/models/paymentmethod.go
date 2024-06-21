package models

import (

)

type PaymentMethodEntity struct {
    Id 			 int        `gorm:"primaryKey;    column:paymentchannel_code;"`
    Name  		 string     `gorm:"not null;      column:paymentchannel_name;"`
}

//make it know table name from database instead of gorm convention
func (p PaymentMethodEntity) TableName() string {
	return "paymentchannel"
}