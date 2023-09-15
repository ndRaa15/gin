package handler

import (
	"gin/internal/domain/service"
)

type Handler struct {
	User       service.UserServiceImpl
	Instrument service.InstrumentServiceImpl
}

func NewHandler(user service.UserServiceImpl, instrument service.InstrumentServiceImpl) *Handler {
	return &Handler{
		User: user,
	}
}
