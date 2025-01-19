package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=postgres password=nothing dbname=learngo port=5433 sslmode=disable"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	fmt.Println("Database connected successfully")
	db = d

}

func GetDB() *gorm.DB {
	return db
}
