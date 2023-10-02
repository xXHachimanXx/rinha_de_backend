package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("Error initializing postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
