package repository

import (
	"context"

	"gin/internal/domain/entity"
)

type StudioRepositoryImpl interface {
	GetAll(ctx context.Context) ([]*entity.Studio, error)
	GetByID(ctx context.Context, id uint) (*entity.Studio, error)
	RentStudio(ctx context.Context, rentStudio *entity.RentStudio) (*entity.RentStudio, error)
	Update(ctx context.Context, studio *entity.Studio) (*entity.Studio, error)
	UpdateStartTime(ctx context.Context, studioID, timeID uint, status bool) error
	UpdateEndTime(ctx context.Context, studioID, timeID uint, status bool) error
}
