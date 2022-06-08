package models

import "time"

type Occupancy struct {
	PoolId    int `json:"organizationUnitId"`
	Capacity  int `json:"maxPersonCount"`
	Occupancy int `json:"personCount"`
	Timestamp time.Time
}
