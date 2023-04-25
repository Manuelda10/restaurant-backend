package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine{
	router := gin.New()
	router.GET("/plates", GetAllPlates)
	router.GET("/plates/:id", GetSinglePlate)
	return router
}