package controllers

import "github.com/gin-gonic/gin"

func Setup(router *gin.RouterGroup) {
	PoolsRegister(router)
}
