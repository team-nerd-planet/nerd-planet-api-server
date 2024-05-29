package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
			AllowOrigins:     []string{"*"},
		})
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		slog.Error(err.Error())
	}

	c.JSON(http.StatusInternalServerError, "")
}
