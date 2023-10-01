package config

import "gorm.io/gorm"

var (
	logger *Logger
	db     *gorm.DB
)

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
