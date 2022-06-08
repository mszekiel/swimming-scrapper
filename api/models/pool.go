package models

type Pool struct {
	Identificator int      `json:"identificator"`
	Name          string   `json:"name"`
	LocationID    int      `json:"-"`
	Location      Location `json:"location"`
}
