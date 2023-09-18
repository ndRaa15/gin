package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Instrument struct {
	ID            uint      `json:"id" gorm:"autoIncreament;primaryKey"`
	Name          string    `json:"name"`
	Owner         string    `json:"owner"`
	ShortDesc     string    `json:"short_desc"`
	Description   string    `json:"description"`
	RentPrice     float64   `json:"rent_price"`
	District      string    `json:"district"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Street        string    `json:"street"`
	Spesification string    `json:"spesification"`
	OwnerNumber   string    `json:"owner_number"`
	Weight        int       `json:"weight"`
	IsBooked      bool      `json:"is_booked"`
	Photo         string    `json:"photo"`
	Rating        float64   `json:"rating"`
	CreateAt      time.Time `json:"-" gorm:"autoCreateTime"`
	UpdateAt      time.Time `json:"-" gorm:"autoUpdateTime"`
}

type RentInstrument struct {
	ID             uint       `json:"id" gorm:"autoIncreament;primaryKey"`
	UserID         uuid.UUID  `json:"user_id" gorm:"type:char(36)"`
	InstrumentID   uint       `json:"-"`
	Instrument     Instrument `json:"-" gorm:"foreignKey:InstrumentID"`
	StartDate      string     `json:"start_date"`
	LengthLoan     uint       `json:"length_loan" binding:"required"`
	Courier        string     `json:"courier" binding:"required"`
	RentCost       float64    `json:"rent_cost" binding:"required"`
	ShippingCost   float64    `json:"shipping_cost" binding:"required"`
	ServiceCost    float64    `json:"service_cost" binding:"required"`
	TotalCost      float64    `json:"total_cost"`
	EstimationTime string     `json:"estimation_time" binding:"required"`
	Status         string     `json:"status" gorm:"default:'PENDING'"`
	Note           string     `json:"note"`
}

type ShippingCost struct {
	ProvinceDestination string `json:"province_destination"`
	CityDestination     string `json:"city_destination"`
}
