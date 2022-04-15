package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://functions.api.ticos-systems.cloud/api/gates/counter?organizationUnitIds=30208", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header = http.Header{
		"Abp-TenantId": []string{"69"},
		"Abp.TenantId": []string{"69"},
		"Origin":       []string{"https://www.swm.de"},
		"Referer":      []string{"https://www.swm.de/"},
	}

	c.AddFunc("*/5 * * * *", func() {
		res, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		var j interface{}
		err = json.NewDecoder(res.Body).Decode(&j)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s", j)
	})

	c.Start()
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
