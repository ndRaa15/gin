package repository

import (
	"context"

	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"gin/internal/infrastructure/mysql"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) repository.UserRepositoryImpl {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *entity.User, ctx context.Context) (*entity.User, error) {
	if err := ur.db.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindByID(id uuid.UUID, ctx context.Context) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Debug().WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *entity.User, ctx context.Context, id uuid.UUID) (*entity.User, error) {
	if err := ur.db.Debug().WithContext(ctx).Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindByEmail(email string, ctx context.Context) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
