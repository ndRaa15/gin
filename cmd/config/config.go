package config

import (
	"fmt"
	"os"
)

func MySQLConfig() string {
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("CONFIG_DB_USER"),
		os.Getenv("CONFIG_DB_PASSWORD"),
		os.Getenv("CONFIG_DB_HOST"),
		os.Getenv("CONFIG_DB_PORT"),
		os.Getenv("CONFIG_DB_NAME"),
	)
	return DataSourceName
}
