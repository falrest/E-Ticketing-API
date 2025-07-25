package controllers

import (
	"e-ticketing-api/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTerminal(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	c.BindJSON(&input)

	_, err := database.DB.Exec("INSERT INTO terminals (name, location) VALUES ($1, $2)", input.Name, input.Location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan terminal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Terminal berhasil dibuat"})
}
