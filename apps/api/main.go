package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	_ "github.com/tsungyu927/bill-breaker/api/docs"
	"github.com/tsungyu927/bill-breaker/api/middlewares"
	"github.com/tsungyu927/bill-breaker/api/routers"
	"github.com/tsungyu927/bill-breaker/api/constants"
	"github.com/tsungyu927/bill-breaker/api/db"
)

// @title Bill Breaker API
// @version 1.0
// @description This is the API of Bill Breaker

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Init database
	db.InitDB()
	defer db.CloseDB()

	// Init gin
	router := gin.Default()
	
	// Middleware
	router.Use(middlewares.CORSMiddleware(constants.APIWhitelist))

	// v1 routers
	v1 := router.Group("api/v1")
	routers.BookRegister(v1)	
	routers.UserRegister(v1)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run
	router.Run("localhost:8080")

	fmt.Println("Hi, I'm getting started")	
}
