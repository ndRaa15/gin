package repository

import (
	"context"

	"gin/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type UserRepositoryImpl interface {
	Create(user *entity.User, ctx context.Context) (*entity.User, error)
	FindByID(id uuid.UUID, ctx context.Context) (*entity.User, error)
	FindByEmail(email string, ctx context.Context) (*entity.User, error)
	Update(user *entity.User, ctx context.Context, id uuid.UUID) (*entity.User, error)
}
