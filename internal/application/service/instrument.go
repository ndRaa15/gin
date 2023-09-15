package service

import (
	"context"
	"os"

	"gin/global/errors"
	"gin/global/time"
	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"gin/internal/infrastructure/rajaongkir"

	"github.com/gofrs/uuid"
)

type InstrumentService struct {
	InstrumentRepository repository.InstrumentRepositoryImpl
}

func NewInstrumentService(instrumentRepository repository.InstrumentRepositoryImpl) *InstrumentService {
	return &InstrumentService{
		InstrumentRepository: instrumentRepository,
	}
}

func (is *InstrumentService) GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error) {
	instruments, err := is.InstrumentRepository.GetAllInstrument(ctx)
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

func (is *InstrumentService) GetByID(ctx context.Context, id uint) (*entity.Instrument, error) {
	instrument, err := is.InstrumentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}

func (is *InstrumentService) RentInstrument(ctx context.Context, id uint, req *entity.RentInstrument, idUser uuid.UUID) (*entity.RentInstrument, error) {
	instrument, err := is.InstrumentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if instrument.IsBooked {
		return nil, errors.ErrInstrumentIsBooked
	}

	totalCost := req.RentCost + req.ShippingCost + req.ServiceCost

	rentInstrument := &entity.RentInstrument{
		UserID:         idUser,
		InstrumentID:   id,
		Courier:        req.Courier,
		StartDate:      time.GenerateDate(),
		LengthLoan:     req.LengthLoan,
		RentCost:       req.RentCost,
		ShippingCost:   req.ShippingCost,
		ServiceCost:    req.ServiceCost,
		TotalCost:      totalCost,
		EstimationTime: req.EstimationTime,
		Note:           req.Note,
	}

	res, err := is.InstrumentRepository.CreateRentInstrument(ctx, rentInstrument)
	if err != nil {
		return nil, err
	}

	instrument.IsBooked = true
	_, err = is.InstrumentRepository.Update(ctx, instrument, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (is *InstrumentService) GetProvince(ctx context.Context, idProvince string) (*rajaongkir.RajaOngkirResponseProvince, error) {
	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))
	province, err := rajaOngkir.GetProvince(idProvince)
	if err != nil {
		return nil, err
	}
	return province, nil
}

func (is *InstrumentService) GetCity(ctx context.Context, idProvince, idCity string) (*rajaongkir.RajaOngkirResponseCity, error) {
	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))
	city, err := rajaOngkir.GetCity(idCity, idProvince)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (is *InstrumentService) GetCost(ctx context.Context, id uint, req *entity.ShippingCost) ([]*rajaongkir.RajaOngkirResponseCost, error) {
	if req.CityDestination == "" || req.ProvinceDestination == "" {
		return nil, errors.ErrBadRequest
	}

	instrument, err := is.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))

	province, err := rajaOngkir.GetProvince("")
	if err != nil {
		return nil, err
	}

	idProvinceOrigin := findProvinceID(instrument.Province, province)

	cityOrigin, err := rajaOngkir.GetCity(idProvinceOrigin, "")
	if err != nil {
		return nil, err
	}
	idCityOrigin := findCityID(instrument.City, cityOrigin)

	idProvinceDestination := findProvinceID(req.ProvinceDestination, province)
	cityDestination, err := rajaOngkir.GetCity(idProvinceDestination, "")
	if err != nil {
		return nil, err
	}
	idCityDestination := findCityID(req.CityDestination, cityDestination)

	cost, err := rajaOngkir.GetCost(idCityOrigin, idCityDestination, instrument.Weight)
	if err != nil {
		return nil, err
	}

	return cost, nil
}

func findCityID(city string, response *rajaongkir.RajaOngkirResponseCity) string {
	var targetID string
	for _, result := range response.RajaOngkir.Results {
		if result.CityName == city {
			targetID = result.CityID
			break
		}
	}
	return targetID
}

func findProvinceID(province string, response *rajaongkir.RajaOngkirResponseProvince) string {
	var targetID string
	for _, result := range response.RajaOngkir.Results {
		if result.Province == province {
			targetID = result.ProvinceID
			break
		}
	}
	return targetID
}
