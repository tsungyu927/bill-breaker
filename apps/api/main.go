package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/tsungyu927/bill-breaker/api/constants"
	"github.com/tsungyu927/bill-breaker/api/db"
	_ "github.com/tsungyu927/bill-breaker/api/docs"
	"github.com/tsungyu927/bill-breaker/api/middlewares"
	"github.com/tsungyu927/bill-breaker/api/routers"
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

	// Start server
	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown ther server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to graceful server shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}

	// Close database connection
	log.Println("Closing database connection...")
	db.CloseDB()

	log.Println("Server exiting")
}
