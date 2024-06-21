package repositories

import (
	"timestamp/models"

	"time"
)

type TimestampRepository interface {
	GetAll() ([]models.TimestampReport, error)
	GetAllByDay(date time.Time) ([]models.TimestampReport, error)
	GetAllByMonth(date time.Time) ([]models.TimestampReport, error)
	GetAllByYear(date time.Time) ([]models.TimestampReport, error)
}