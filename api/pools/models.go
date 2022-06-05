package pools

import (
	"mszekiel/swimming-scrapper/common"

	"gorm.io/gorm"
)

type Location struct {
	ID  int
	Lat float64 `gorm:"type:decimal(10,8)"`
	Lon float64 `gorm:"type:decimal(11,8)"`
}

type Pool struct {
	gorm.Model
	Identificator int
	Name          string
	LocationID    int
	Location      Location
}

func GetAllPools() ([]Pool, int64, error) {
	db := common.GetConnection()

	var pools []Pool
	var count int64

	err := db.Find(&pools).Count(&count).Error

	return pools, count, err
}

func SaveOne(data interface{}) error {
	db := common.GetConnection()
	err := db.Save(data).Error

	return err
}
