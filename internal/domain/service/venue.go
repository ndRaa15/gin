package service

import (
	"context"

	"gin/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type VenueServiceImpl interface {
	GetAllVenue(ctx context.Context) ([]*entity.Venue, error)
	GetVenueByID(ctx context.Context, id uint) (*entity.Venue, error)
	RentVenue(ctx context.Context, userID uuid.UUID, venueDayID uint) (*entity.ApplyVenue, error)
	GetListApplyVenue(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error)
}
