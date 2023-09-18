package service

import (
	"context"
	"errors"

	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"github.com/gofrs/uuid"
)

type StudioService struct {
	StudioRepository repository.StudioRepositoryImpl
}

func NewStudioService(studioRepository repository.StudioRepositoryImpl) *StudioService {
	return &StudioService{
		StudioRepository: studioRepository,
	}
}

func (ss *StudioService) GetAll(ctx context.Context) ([]*entity.Studio, error) {
	studios, err := ss.StudioRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (ss *StudioService) GetByID(ctx context.Context, id uint) (*entity.Studio, error) {
	studio, err := ss.StudioRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (ss *StudioService) RentStudio(ctx context.Context, studioID uint, userID uuid.UUID, req *entity.RentStudio) (*entity.RentStudio, error) {
	studio, err := ss.StudioRepository.GetByID(ctx, studioID)
	if err != nil {
		return nil, err
	}

	var indexStartTime int
	var indexEndTime int

	for i := 0; i < len(studio.StartTime); i++ {
		if studio.StartTime[i].Time.Time == req.StartTime && studio.StartTime[i].IsAvailable {
			indexStartTime = i
		}
		if studio.EndTime[i].Time.Time == req.EndTime && studio.EndTime[i].IsAvailable {
			indexEndTime = i
		}
	}

	if indexStartTime >= indexEndTime {
		return nil, errors.New("INVALID TIME")
	} else if indexStartTime == 0 || indexEndTime == 0 {
		return nil, errors.New("TIME NOT AVAILABLE")
	}

	// Make start time and end time to be slice
	for i := indexStartTime; i <= indexEndTime; i++ {
		if err := ss.StudioRepository.UpdateStartTime(ctx, studioID, studio.StartTime[i].TimeID, false); err != nil {
			return nil, err
		}
		if err := ss.StudioRepository.UpdateEndTime(ctx, studioID, studio.EndTime[i].TimeID, false); err != nil {
			return nil, err
		}
	}

	_, err = ss.StudioRepository.Update(ctx, studio)
	if err != nil {
		return nil, err
	}

	rentStudio := &entity.RentStudio{
		UserID:    userID,
		StudioID:  studioID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		TotalHour: uint(indexEndTime) - uint(indexStartTime) + 1,
		TotalCost: (studio.PricePerHour * float64(uint(indexEndTime)-uint(indexStartTime)+1)) + req.ServiceCost,
	}

	res, err := ss.StudioRepository.RentStudio(ctx, rentStudio)
	if err != nil {
		return nil, err
	}

	return res, nil
}
