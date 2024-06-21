package mock_services

import (
	"github.com/stretchr/testify/mock"
	"time"

	"timestamp/models"
)

type timestampRepositoryMock struct {
	mock.Mock
}

func NewTimestampServiceMock() *timestampRepositoryMock {
	return &timestampRepositoryMock{}
}

func (m *timestampRepositoryMock) GetTimestamps() ([]models.TimestampReport, error) {
	args := m.Called()
	if args.Get(0) != nil {
        return args.Get(0).([]models.TimestampReport), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *timestampRepositoryMock) GetTimestampsByDay(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	if args.Get(0) != nil {
        return args.Get(0).([]models.TimestampReport), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *timestampRepositoryMock) GetTimestampsByMonth(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	if args.Get(0) != nil {
        return args.Get(0).([]models.TimestampReport), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *timestampRepositoryMock) GetTimestampsByYear(date time.Time) ([]models.TimestampReport, error) {
	args := m.Called(date)
	if args.Get(0) != nil {
        return args.Get(0).([]models.TimestampReport), args.Error(1)
    }
    return nil, args.Error(1)
}