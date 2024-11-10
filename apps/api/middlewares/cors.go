package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(whitelist []string) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     whitelist,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow credentials (cookies, auth headers)
		MaxAge:           12 * time.Hour, // Cache duration for preflight response
	}

	return cors.New(corsConfig)
}
