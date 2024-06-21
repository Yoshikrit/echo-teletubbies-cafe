package models

import (
	"time"
)

type TimestampEntity struct {
	Id               int         `gorm:"primaryKey;  autoIncrement"`
    UserId  		 int         `gorm:"not null;    column:timestamp_user_code;"`
	LoginAt          time.Time   `gorm:"not null;    column:timestamp_login_at;"`
	LogoutAt         time.Time   `gorm:"             column:timestamp_logout_at;"`
	Hour             int         `gorm:"             column:timestamp_hour;"`
}

//make it know table name from database instead of gorm convention
func (p TimestampEntity) TableName() string {
	return "timestamp"
}

type Timestamp struct {
	Seq          int        `json:"Timestamp_Seq"`
    UserId     	 int        `json:"Timestamp_UserId"`
	LoginAt      time.Time 	`json:"Timestamp_LoginAt"`  
	LogoutAt     time.Time 	`json:"Timestamp_LogoutAt"`  
	Hour         int        `json:"Timestamp_Hour"`
}

type TimestampReport struct {
	Seq          int        `json:"Timestamp_Seq"`
    UserName     string     `json:"Timestamp_UserName"`
	LoginAt      time.Time 	`json:"Timestamp_LoginAt"`  
	LogoutAt     time.Time 	`json:"Timestamp_LogoutAt"`  
	Hour         int        `json:"Timestamp_Hour"`
}

