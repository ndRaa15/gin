package handler

import (
	"gin/internal/domain/service"
)

type Handler struct {
	User       service.UserServiceImpl
	Instrument service.InstrumentServiceImpl
	Venue      service.VenueServiceImpl
	Studio     service.StudioServiceImpl
}

func NewHandler(user service.UserServiceImpl, instrument service.InstrumentServiceImpl, venue service.VenueServiceImpl, studio service.StudioServiceImpl) *Handler {
	return &Handler{
		User:       user,
		Instrument: instrument,
		Venue:      venue,
		Studio:     studio,
	}
}
