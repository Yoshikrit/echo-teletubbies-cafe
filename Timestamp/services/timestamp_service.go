package services

import (
	"timestamp/models"
	"timestamp/repositories"
	// "timestamp/utils/errs"
	"timestamp/utils/logs"

	"time"
)

type timestampService struct {
	timestampRepo repositories.TimestampRepository
}

func NewTimestampService(timestampRepo repositories.TimestampRepository) TimestampService {
	return timestampService{timestampRepo: timestampRepo}
}

func (s timestampService) GetTimestamps() ([]models.TimestampReport, error) {
	timestampsReport, err := s.timestampRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get Timestamps Successfully")
	return timestampsReport, nil
}

func (s timestampService) GetTimestampsByDay(dateReq time.Time) ([]models.TimestampReport, error) {
	timestampsReport, err := s.timestampRepo.GetAllByDay(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get Timestamps By Day Successfully")
	return timestampsReport, nil
}

func (s timestampService) GetTimestampsByMonth(dateReq time.Time) ([]models.TimestampReport, error) {
	timestampsReport, err := s.timestampRepo.GetAllByMonth(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get Timestamps By Month Successfully")
	return timestampsReport, nil
}

func (s timestampService) GetTimestampsByYear(dateReq time.Time) ([]models.TimestampReport, error) {
	timestampsReport, err := s.timestampRepo.GetAllByYear(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get Timestamps By Year Successfully")
	return timestampsReport, nil
}