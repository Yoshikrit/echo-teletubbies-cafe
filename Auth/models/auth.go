package models

import (
)

type UserLogin struct {
    Email 		 string     `json:"User_Email"         validate:"required"`
    Password 	 string     `json:"User_Password"      validate:"required"`
}

type UserClaim struct {
    Id           int        
    Name 		 string     
    Role 	     string 
}

type Response struct {
    Id           int        `json:"User_Id"`
    Name 		 string     `json:"User_Name"`
    Role 	     string     `json:"User_Role"`
    JWT 	     string     `json:"JWT"`
}
