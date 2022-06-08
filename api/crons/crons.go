package crons

import (
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func initCrons(c *cron.Cron) {
	LogOccupancy()
	c.AddFunc("*/5 * * * *", LogOccupancy)
}

func Setup() {
	Cron = cron.New()

	initCrons(Cron)

	Cron.Start()
}
