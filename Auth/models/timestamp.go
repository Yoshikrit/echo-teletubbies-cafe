package models

import (
	"time"
)

type TimestampEntity struct {
	Id               int         `gorm:"primary_key; column:timestamp_code;"`
    UserId  		 int         `gorm:"not null;    column:timestamp_user_code;"`
	LoginAt          time.Time   `gorm:"not null;    column:timestamp_login_at;"`
	LogoutAt         time.Time   `gorm:"             column:timestamp_logout_at;"`
	Hour             int         `gorm:"             column:timestamp_hour;"`
}

//make it know table name from database instead of gorm convention
func (p TimestampEntity) TableName() string {
	return "timestamp"
}
