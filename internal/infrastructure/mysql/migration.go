package mysql

import (
	"log"

	"gin/global/errors"
	"gin/internal/domain/entity"
)

func Migration(db *DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		log.Fatalf("[musiku-postgresql] failed to migrate musiku database : %v\n", err)
		return errors.ErrMigrateDatabase
	}
	return nil
}
