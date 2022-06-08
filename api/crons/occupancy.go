package crons

import (
	"fmt"
	"mszekiel/swimming-scrapper/common"
	"mszekiel/swimming-scrapper/models"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

func getOccupancyData() []models.Occupancy {
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

	var occupancyData []models.Occupancy
	err = json.NewDecoder(res.Body).Decode(&occupancyData)

	if err != nil {
		log.Fatalln(err)
	}

	currentTime := time.Now().UTC()
	for i := range occupancyData {

		occupancyData[i].Timestamp = currentTime
	}

	return occupancyData
}

func LogOccupancy() {
	fmt.Println("Loggin Occupancy", time.Now().UTC())
	db := common.GetConnection()
	data := getOccupancyData()
	db.Create(data)
}
