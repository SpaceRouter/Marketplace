package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/marketplace/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/tea", controllers.GetTea)

	return router

}
