package controllers

import (
	"e-ticketing-api/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"os"
	"net/http"
)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BindJSON(&input)

	row := database.DB.QueryRow("SELECT id FROM users WHERE username=$1 AND password=$2", input.Username, input.Password)
	var id int
	err := row.Scan(&id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login gagal"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
