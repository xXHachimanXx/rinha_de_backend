package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := "host=localhost user=admin password=admin dbname=persondatabase port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Postgres opening error: %v", err)
		return nil, err
	}

	return db, nil
}
