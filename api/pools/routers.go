package pools

import (
	"errors"
	"mszekiel/swimming-scrapper/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PoolsRegister(router *gin.RouterGroup) {
}

func PoolsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", PoolsList)
	router.POST("/", PoolCreate)
}

func PoolsList(c *gin.Context) {
	poolModels, poolCount, err := GetAllPools()

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("pools", errors.New("invalid param")))
		return
	}

	serializer := PoolsSerializer{c, poolModels}

	c.JSON(http.StatusOK, gin.H{"count": poolCount, "pools": serializer.Response()})
}

func PoolCreate(c *gin.Context) {
	poolValidator := NewPoolValidator()

	if err := poolValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&poolValidator.poolModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("pools", err))
		return
	}

	serializer := PoolSerializer{c, poolValidator.poolModel}

	c.JSON(http.StatusCreated, gin.H{"pool": serializer.Response()})
}
