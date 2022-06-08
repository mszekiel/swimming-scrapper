package common

import (
	"mszekiel/swimming-scrapper/models"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Migrate() {
	db := GetConnection()

	db.AutoMigrate(&models.Pool{})
	db.AutoMigrate(&models.Occupancy{})
}

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("PG_DSN")

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = _db

	Migrate()

	return DB
}

func GetConnection() *gorm.DB {
	return DB
}
