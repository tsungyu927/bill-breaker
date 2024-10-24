package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/routers"
)

func main() {
	// Init gin
	router := gin.Default()
	
	// Middleware
	// TODO: add CORS, Auth middlewares

	// v1
	v1 := router.Group("api/v1")
	routers.BookRegister(v1)	


	// Run
	router.Run("localhost:8080")

	fmt.Println("Hi, I'm getting started")	
}
