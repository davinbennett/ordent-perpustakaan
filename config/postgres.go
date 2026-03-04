package config

import (
	"fmt"
	"log"
	"ordentperpustakaan/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresDB *gorm.DB

func InitPostgres() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	PostgresDB = db

	PostgresDB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	return nil
}

func AutoMigrate() error {
	return PostgresDB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Book{},
		&models.Loan{},
		&models.Reservation{},
	)
}