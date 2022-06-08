package main

import (
	"mszekiel/swimming-scrapper/common"
	"mszekiel/swimming-scrapper/controllers"
	"mszekiel/swimming-scrapper/crons"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	common.InitCache()
	common.ConnectDatabase()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://swimming.up.railway.app", "http://localhost:3000"},
		AllowMethods: []string{"GET"},
	}))

	v1 := router.Group("/api")

	controllers.Setup(v1)
	crons.Setup()

	router.Run()
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
