package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_DSN")), &gorm.Config{})
}
