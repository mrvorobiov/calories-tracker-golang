package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
			os.Getenv("PG_HOST"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_PWD"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_PORT"))))

	if err != nil {
		return nil, err
	}

	driver, err := db.DB()

	if err != nil {
		return nil, err
	}

	if err := driver.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
