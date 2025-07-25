package main

import (
	"e-ticketing-api/controllers"
	"e-ticketing-api/database"
	"e-ticketing-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectDB()

	r := gin.Default()

	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/terminals", controllers.CreateTerminal)

	r.Run(":8080")
}
