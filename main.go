package main

import (
	"encoding/json"
	"log"
	"mszekiel/swimming-scrapper/database"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
)

type Pool struct {
	PoolId    uint `json:"organizationUnitId"`
	Capacity  uint `json:"maxPersonCount"`
	Occupancy uint `json:"personCount"`
	Timestamp time.Time
}

func getPoolData() []Pool {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://functions.api.ticos-systems.cloud/api/gates/counter?organizationUnitIds=30208&organizationUnitIds=30194&organizationUnitIds=30195&organizationUnitIds=30190&organizationUnitIds=129&organizationUnitIds=30208&organizationUnitIds=30197&organizationUnitIds=30184&organizationUnitIds=30182&organizationUnitIds=30187&organizationUnitIds=30199&organizationUnitIds=30204&organizationUnitIds=30207&organizationUnitIds=30188&organizationUnitIds=30200&organizationUnitIds=30206&organizationUnitIds=30185&organizationUnitIds=30203&organizationUnitIds=30191&organizationUnitIds=30201", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header = http.Header{
		"Abp-TenantId": []string{"69"},
		"Abp.TenantId": []string{"69"},
		"Origin":       []string{"https://www.swm.de"},
		"Referer":      []string{"https://www.swm.de/"},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var poolData []Pool
	err = json.NewDecoder(res.Body).Decode(&poolData)

	if err != nil {
		log.Fatalln(err)
	}

	for i := range poolData {
		poolData[i].Timestamp = time.Now().UTC()
	}

	return poolData
}

func logPools() {
	db := database.GetConnection()
	data := getPoolData()

	for _, pool := range data {
		db.Create(pool)
	}
}

func initDatabase() {
	db := database.GetConnection()

	db.AutoMigrate(&Pool{})

}

func main() {
	c := cron.New()

	initDatabase()
	logPools()

	c.AddFunc("*/5 * * * *", logPools)
	c.Start()

	select {}
}

// Swimming pools
// 1.  30194 -> Bad Forstenrieder Park - Hallenbad
// 2.  30195 -> Bad Giesing-Harlaching - Hallenbad
// 3.  30190 -> Cosimawellenbad - Hallenbad
// 4.  129 -> Dante-Winter-Warmfreibad
// 5.  30208 -> Michaelibad - Hallenbad
// 6.  30197 -> Müller‘sches Volksbad - Hallenbad
// 7.  30184 -> Nordbad - Hallenbad & Fitness
// 8.  30182 -> Olympia-Schwimmhalle
// 9.  30187 -> Südbad - Hallenbad
// 10. 30199 -> Westbad - Hallenbad
//
// Saunas
// 1.  30204 -> Müller‘sches Volksbad - Sauna
// 2.  30207 -> Westbad - Sauna
// 3.  30188 -> Südbad - Sauna
// 4.  30200 -> Dantebad - Sauna
// 5.  30206 -> Prinzregentenbad - Sauna
// 6.  30185 -> Nordbad - Sauna
// 7.  30203 -> Michaelibad - Sauna
// 8.  30191 -> Cosimawellenbad - Sauna
// 9.  30201 -> Bad Forstenrieder Park - Sauna
