package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/marketplace/controllers"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/tea", controllers.GetTea)

	v1 := router.Group("v1")
	{
		stack := controllers.StackController{
			DB: db,
		}
		v1.GET("id/:id", stack.GetStackById)
	}

	return router
}
