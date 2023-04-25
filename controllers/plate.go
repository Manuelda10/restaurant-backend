package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restaurant-backend/models"
)

func GetAllPlates (c *gin.Context){
	//w.Header().Set("Content-Type", "application/json")

	var plates []models.Plate
	models.DB.Find(&plates)

	c.IndentedJSON(http.StatusOK, plates)
	//json.NewEncoder(w).Encode(plates)
}

func GetSinglePlate (c *gin.Context){
	var plate models.Plate

	if err := models.DB.Where("id=?", c.Param("id")).First(&plate).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Registro no encontrado"})
		return
	}
	c.IndentedJSON(http.StatusOK, plate)
}


