package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/middlewares"
	"github.com/tsungyu927/bill-breaker/api/routers"
	"github.com/tsungyu927/bill-breaker/api/constants"
)

func main() {
	// Init gin
	router := gin.Default()
	
	// Middleware
	router.Use(middlewares.CORSMiddleware(constants.APIWhitelist))

	// v1 routers
	v1 := router.Group("api/v1")
	routers.BookRegister(v1)	


	// Run
	router.Run("localhost:8080")

	fmt.Println("Hi, I'm getting started")	
}
