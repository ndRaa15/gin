package handler

import (
	"gin/internal/domain/service"
)

type Handler struct {
	User service.UserServiceImpl
}

func NewHandler(user service.UserServiceImpl) *Handler {
	return &Handler{
		User: user,
	}
}
