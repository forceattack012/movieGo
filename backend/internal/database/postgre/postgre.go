package postgre

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

type PostgreConfig struct {
	host         string
	user         string
	password     string
	databaseName string
	port         string
}

func BuildConfig(h string, u string, p string, db string, port string) *PostgreConfig {
	return &PostgreConfig{
		host:         h,
		user:         u,
		password:     p,
		databaseName: db,
		port:         port,
	}
}

func Connect(config *PostgreConfig) {
	pgConn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", config.host, config.user, config.password, config.port)
	db, err := gorm.Open(postgres.Open(pgConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	Database = db
}
