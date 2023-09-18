package repository

import (
	"fmt"

	"gin/internal/domain/entity"
	"gin/internal/infrastructure/mysql"
)

func SeedTime(db *mysql.DB) error {
	var dummyData []entity.Time

	for i := 1; i <= 24; i++ {
		hour := fmt.Sprintf("%02d", i)
		time := hour + ":00"
		entity := entity.Time{
			ID:   uint(i),
			Time: time,
		}
		dummyData = append(dummyData, entity)
	}

	if err := db.Create(&dummyData).Error; err != nil {
		return err
	}
	return nil
}
