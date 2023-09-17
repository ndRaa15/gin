package service

import (
	"context"

	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"gin/internal/domain/service"
	"github.com/gofrs/uuid"
)

type VenueService struct {
	VenueRepository repository.VenueRepositoryImpl
}

func NewVenueService(VenueRepository repository.VenueRepositoryImpl) service.VenueServiceImpl {
	return &VenueService{VenueRepository}
}

func (vs *VenueService) GetAllVenue(ctx context.Context) ([]*entity.Venue, error) {
	venues, err := vs.VenueRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return venues, nil
}

func (vs *VenueService) GetVenueByID(ctx context.Context, id uint) (*entity.Venue, error) {
	venue, err := vs.VenueRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return venue, nil
}

func (vs *VenueService) RentVenue(ctx context.Context, userID uuid.UUID, venueDayID uint) (*entity.ApplyVenue, error) {
	venueDay, err := vs.VenueRepository.GetVenueDayByID(ctx, venueDayID)
	if err != nil {
		return nil, err
	}

	if venueDay.Status == "BOOKED" {
		return nil, err
	}

	applyVenue := entity.ApplyVenue{
		UserID:     userID,
		VenueDayID: venueDayID,
		VenueDay:   *venueDay,
	}

	res, err := vs.VenueRepository.CreateApplyVenue(ctx, &applyVenue)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VenueService) GetListApplyVenue(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error) {
	applyVenues, err := vs.VenueRepository.GetApplyVenueByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return applyVenues, nil
}
