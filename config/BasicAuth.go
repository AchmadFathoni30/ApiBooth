package config

import (
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"Dispenda": "*#Dispenda*#2024@",
	})
}
