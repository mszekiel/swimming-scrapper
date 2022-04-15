package main

import (
	"log"
	"net/http"
)

func main() {
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

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	println(res)
}
