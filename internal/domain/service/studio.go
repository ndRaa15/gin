package service

import (
	"context"

	"gin/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type StudioServiceImpl interface {
	GetAll(ctx context.Context) ([]*entity.Studio, error)
	GetByID(ctx context.Context, id uint) (*entity.Studio, error)
	RentStudio(ctx context.Context, studioID uint, userID uuid.UUID, req *entity.RentStudio) (*entity.RentStudio, error)
}
