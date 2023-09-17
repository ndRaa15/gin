package repository

import (
	"gin/internal/domain/entity"
	"gin/internal/infrastructure/mysql"
)

func SeedDays(db *mysql.DB) error {
	var days []entity.Day
	monday := entity.Day{
		ID:   1,
		Name: "Monday",
	}

	tuesday := entity.Day{
		ID:   2,
		Name: "Tuesday",
	}

	wednesday := entity.Day{
		ID:   3,
		Name: "Wednesday",
	}

	thursday := entity.Day{
		ID:   4,
		Name: "Thursday",
	}

	friday := entity.Day{
		ID:   5,
		Name: "Friday",
	}

	saturday := entity.Day{
		ID:   6,
		Name: "Saturday",
	}

	sunday := entity.Day{
		ID:   7,
		Name: "Sunday",
	}

	days = append(days, monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	for _, day := range days {
		if err := db.Create(&day).Error; err != nil {
			return err
		}
	}
	return nil
}
