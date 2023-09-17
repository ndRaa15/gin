package repository

import (
	"context"

	"gin/internal/domain/entity"
	"github.com/gofrs/uuid"
)

type VenueRepositoryImpl interface {
	GetAll(ctx context.Context) ([]*entity.Venue, error)
	GetByID(ctx context.Context, id uint) (*entity.Venue, error)
	GetVenueDayByID(ctx context.Context, venueDayID uint) (*entity.VenueDay, error)
	CreateApplyVenue(ctx context.Context, applyVenue *entity.ApplyVenue) (*entity.ApplyVenue, error)
	GetApplyVenueByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error)
}
