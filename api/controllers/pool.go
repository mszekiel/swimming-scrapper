package controllers

import (
	"mszekiel/swimming-scrapper/common"
	"mszekiel/swimming-scrapper/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

var validate *validator.Validate

func PoolsRegister(router *gin.RouterGroup) {
	router.GET("/pools", FindPools)
	// router.POST("/pools", CreatePool)

	validate = validator.New()
}

func FindPools(c *gin.Context) {
	if _pools, err := common.Cache.Get("pools"); err == nil {
		var cacheData map[string]interface{}

		json.Unmarshal(_pools, &cacheData)
		c.JSON(http.StatusOK, cacheData)
		return
	}

	var pools []models.Pool
	var count int64

	common.DB.Preload("Location").Find(&pools).Count(&count)
	data, _ := json.Marshal(gin.H{"data": pools, "count": count})

	common.Cache.Set("pools", data)

	c.JSON(http.StatusOK, gin.H{"data": pools, "count": count})
}

type CreatePoolInput struct {
	Name          string          `json:"name" binding:"required" validate:"required"`
	Identificator int             `json:"identificator" binding:"required" validate:"required"`
	Location      models.Location `json:"location" binding:"required" validate:"required"`
}

func CreatePool(c *gin.Context) {
	var input CreatePoolInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pool := models.Pool{Name: input.Name, Identificator: input.Identificator, Location: input.Location}

	common.DB.Create(&pool)

	c.JSON(http.StatusOK, gin.H{"pool": pool})

}
