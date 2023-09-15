package service

import (
	"context"
	"gin/internal/domain/entity"
	"gin/internal/infrastructure/rajaongkir"

	"github.com/gofrs/uuid"
)

type InstrumentServiceImpl interface {
	GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error)
	GetByID(ctx context.Context, id uint) (*entity.Instrument, error)
	RentInstrument(ctx context.Context, id uint, req *entity.RentInstrument, idUser uuid.UUID) (*entity.RentInstrument, error)
	GetProvince(ctx context.Context, idProvince string) (*rajaongkir.RajaOngkirResponseProvince, error)
	GetCity(ctx context.Context, idProvince, idCity string) (*rajaongkir.RajaOngkirResponseCity, error)
	GetCost(ctx context.Context, id uint, req *entity.ShippingCost) ([]*rajaongkir.RajaOngkirResponseCost, error)
}
