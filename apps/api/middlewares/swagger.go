package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SwaggerAuthMiddleware() gin.HandlerFunc {
	username := os.Getenv("SWAGGER_USER")
	password := os.Getenv("SWAGGER_PASSWORD")

	if username == "" || password == "" {
		panic("SWAGGER_USER or SWAGGER_PASSWORD is not set")
	}

	authUsers := gin.Accounts{
		username: password,
	}

	return gin.BasicAuth(authUsers)
}
