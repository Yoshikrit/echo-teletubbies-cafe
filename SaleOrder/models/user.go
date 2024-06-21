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

type User struct {
    Id     		 int    	`json:"User_Id"`
	Role_Id      int 	    `json:"Role_Id"`
    FName 		 string     `json:"User_FName"`
    LName 		 string     `json:"User_LName"`
    Email 		 string     `json:"User_Email"`
    Sex 		 string     `json:"User_Sex"`
    TelNo 		 string     `json:"User_TelNo"`
    Salary 		 float64    `json:"User_Salary"`
    Address 	 string     `json:"User_Address"`
    WorkStatus   string     `json:"User_WorkStatus"`
    BirthDate 	 time.Time  `json:"User_BirthDate"`
}

