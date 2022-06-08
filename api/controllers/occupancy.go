package controllers

import (
	"mszekiel/swimming-scrapper/common"
	"mszekiel/swimming-scrapper/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OccupancyRegister(router *gin.RouterGroup) {
	router.GET("/occupancy", FindOccupancy)
}

type FincOccupancyResponse struct {
	models.Pool
	CurrentOccupancy int `json:"currentOccupancy"`
	AverageOccupancy int `json:"averageOccupancy"`
	Capacity         int `json:"capacity"`
}

func FindOccupancy(c *gin.Context) {
	var response FincOccupancyResponse
	var pool models.Pool

	if err := common.DB.Where("identificator = ?", c.Query("id")).First(&pool).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pool not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
