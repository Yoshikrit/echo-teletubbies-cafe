package configs

import (
	"fmt"
	"os"
	"context"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

type DBConfig struct {
    Host     string
    Port     int
    User     string
    DBName   string
    Password string
    TimeZone string
}

// convert string to int
func portConv(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==============================\n", sql)
}

func DatabaseInit() {
    host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := portConv(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Bangkok", host, user, password, dbName, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: &SqlLogger{}, // check sql with log
	})
    if err != nil {
        panic("Failed to connect to database")
    }
    
}

func GetDB() *gorm.DB {
    return DB
}