package models

type Location struct {
	ID  int     `json:"-"`
	Lat float64 `gorm:"type:decimal(9,7)" json:"lat"`
	Lon float64 `gorm:"type:decimal(10,7)" json:"lon"`
}
