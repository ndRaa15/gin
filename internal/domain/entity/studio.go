package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Studio struct {
	ID           uint        `json:"id" gorm:"autoIncreament;primaryKey"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Description  string      `json:"description"`
	PricePerHour float64     `json:"price_per_hour"`
	OpenHour     string      `json:"open_hour"`
	IsAvailable  bool        `json:"is_available" gorm:"default:true"`
	Phone        string      `json:"phone"`
	Photo        string      `json:"photo"`
	Rating       float64     `json:"rating"`
	StartTime    []StartTime `json:"start_time" gorm:"foreignKey:StudioID"`
	EndTime      []EndTime   `json:"end_time" gorm:"foreignKey:StudioID"`
	CreateAt     time.Time   `json:"-" gorm:"autoCreateTime"`
	UpdateAt     time.Time   `json:"-" gorm:"autoUpdateTime"`
}

type StartTime struct {
	StudioID    uint `json:"studio_id" gorm:"primaryKey"`
	TimeID      uint `json:"time_id" gorm:"primaryKey"`
	Time        Time `json:"time" gorm:"foreignKey:TimeID"`
	IsAvailable bool `json:"isAvailable" gorm:"default:true"`
}

type EndTime struct {
	StudioID    uint `json:"studio_id" gorm:"primaryKey"`
	TimeID      uint `json:"time_id" gorm:"primaryKey"`
	Time        Time `json:"time" gorm:"foreignKey:TimeID"`
	IsAvailable bool `json:"isAvailable" gorm:"default:true"`
}

type RentStudio struct {
	ID          uint      `json:"id" gorm:"autoIncreament;primaryKey"`
	StudioID    uint      `json:"studio_id"`
	UserID      uuid.UUID `json:"user_id"`
	StartTime   string    `json:"start_time" binding:"required"`
	EndTime     string    `json:"end_time" binding:"required"`
	ServiceCost float64   `json:"service_cost" binding:"required"`
	TotalHour   uint      `json:"total_hour"`
	TotalCost   float64   `json:"total_cost"`
	Status      string    `json:"status" sql:"type:ENUM('PENDING', 'BOOKED', 'REJECTED')" gorm:"default:'PENDING'"`
}
