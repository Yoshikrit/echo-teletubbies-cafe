package models

import (
    "time"
)

type UserEntity struct {
    Id 			 int         `gorm:"primaryKey; column:user_code;"`
    Role_Id      int 		 `gorm:"not null;   column:user_role_code;"`
    FName 		 string      `gorm:"            column:user_fname;         size:40;"`
    LName 		 string      `gorm:"            column:user_lname;         size:40;"`
    Email 		 string      `gorm:"not null;   column:user_email;         size:50;"`
    Password 	 string      `gorm:"not null;   column:user_password;"`
    Sex		     string      `gorm:"not null;   column:user_sex;           size:1;"`
    TelNo		 string      `gorm:"            column:user_telno;         size:10;"`
    Salary		 float64     `gorm:"not null;   column:user_salary;        type:decimal(12,2);"`
    Address 	 string      `gorm:"            column:user_address;"`   
    WorkStatus 	 string      `gorm:"not null;   column:user_workstatus;    size:1;"`  
    BirthDate 	 time.Time   `gorm:"not null;   column:user_birthdate;"`  
}

//make it know table name from database instead of gorm convention
func (p UserEntity) TableName() string {
	return "user"
}