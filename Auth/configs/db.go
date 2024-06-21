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

func portConv(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==============================\n", sql)
}

func DatabaseInit() {
	cfg := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	port, err := portConv(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Failed to convert DB_PORT to int")
	}
	cfg.Port = port

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Bangkok", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: &SqlLogger{}, // Check SQL with log
	})
	if err != nil {
		panic("Failed to connect to the database")
	}
}

func GetDB() *gorm.DB {
	return DB
}
