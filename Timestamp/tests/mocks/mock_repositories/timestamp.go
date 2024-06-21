package mock_repositories

import (
	"time"
	"github.com/stretchr/testify/mock"

	"timestamp/models"
)

type timestampRepositoryMock struct {
	mock.Mock
}

func NewTimestampRepositoryMock() *timestampRepositoryMock {
	return &timestampRepositoryMock{}
}

func (m *timestampRepositoryMock) GetAll() ([]models.TimestampReport, error) {
	args := m.Called()
	return args.Get(0).([]models.TimestampReport), args.Error(1)
}

func (m *timestampRepositoryMock) GetAllByDay(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	return args.Get(0).([]models.TimestampReport), args.Error(1)
}

func (m *timestampRepositoryMock) GetAllByMonth(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	return args.Get(0).([]models.TimestampReport), args.Error(1)
}

func (m *timestampRepositoryMock) GetAllByYear(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	return args.Get(0).([]models.TimestampReport), args.Error(1)
}