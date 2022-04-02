package postgre

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

type PostgreConfig struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         string
}

func BuildConfig() *PostgreConfig {
	return &PostgreConfig{
		Host: "",
	}
}

func Connect(config *PostgreConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Host), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
