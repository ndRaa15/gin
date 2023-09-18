package mysql

import (
	"log"

	"gin/global/errors"
	"gin/internal/domain/entity"
)

func Migration(db *DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Instrument{},
		&entity.RentInstrument{},
		&entity.Venue{},
		&entity.VenueDay{},
		&entity.ApplyVenue{},
		&entity.Day{},
		&entity.Studio{},
		&entity.RentStudio{},
		&entity.Time{},
		&entity.StartTime{},
		&entity.EndTime{},
	); err != nil {
		log.Fatalf("[musiku-postgresql] failed to migrate musiku database : %v\n", err)
		return errors.ErrMigrateDatabase
	}
	return nil
}
