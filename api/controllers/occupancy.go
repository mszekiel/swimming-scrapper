package controllers

import "github.com/gin-gonic/gin"

func OccupancyRegister(router *gin.RouterGroup) {
	router.GET("/occupancy", FindOccupancy)
}

func FindOccupancy(c *gin.Context) {

}
