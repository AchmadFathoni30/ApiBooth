package router

import (
	"ApiBooth/config"
	"ApiBooth/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authorized := r.Group("/Booth", config.BasicAuthMiddleware())
	{
		authorized.GET("/BoothSlamet", controller.GetBoothSlametRiyadi)
	}
	r.Group("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "OK"})
	})
	return r
}
