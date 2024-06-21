package services

import (
	"timestamp/models"

	"time"
)

type TimestampService interface {
	GetTimestamps() ([]models.TimestampReport, error)
	GetTimestampsByDay(time.Time) ([]models.TimestampReport, error)
	GetTimestampsByMonth(time.Time) ([]models.TimestampReport, error)
	GetTimestampsByYear(time.Time) ([]models.TimestampReport, error)
}

