package router

import (
	"main/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "OK"})
	})

	krakenGroup := r.Group("/kraken")
	krakenGroup.GET("/health", controllers.KrakenGetInfoHealth)
	krakenGroup.GET("/balance", controllers.KrakenGetBalance)
	krakenGroup.GET("/balanceExtend", controllers.KrakenGetBalanceExtend)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	return r
}