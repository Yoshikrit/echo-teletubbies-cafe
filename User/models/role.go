package models

import (

)

type RoleEntity struct {
    Id 			 int         `gorm:"primaryKey; column:role_code;"`
    Name  		 string      `gorm:"not null;   column:role_name;    size:40;"`
}

//make it know table name from database instead of gorm convention
func (p RoleEntity) TableName() string {
	return "role"
}

type Role struct {
    Id     		 int    	`json:"Role_Id"`
	Name         string 	`json:"Role_Name"`
}

type RoleCreate struct {
    Id     		 int    	`json:"Role_Id"          validate:"required,gt=0"`
	Name         string 	`json:"Role_Name"        validate:"required"`
}

type RoleUpdate struct {
    Name 		 string     `json:"Role_Name"        validate:"required"`
}