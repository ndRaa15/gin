package service

import (
	"context"
	"mime/multipart"

	"gin/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type UserServiceImpl interface {
	Register(req *entity.UserRegister, ctx context.Context) (*entity.User, error)
	VerifyAccount(id uuid.UUID, ctx context.Context) (*entity.User, error)
	Login(req *entity.UserLogin, ctx context.Context) (*entity.ResponseLogin, error)
	UploadPhoto(file *multipart.FileHeader, id uuid.UUID, ctx context.Context) (*entity.User, error)
	UpdateUser(req *entity.UserUpdate, ctx context.Context, id uuid.UUID) (*entity.User, error)
}
