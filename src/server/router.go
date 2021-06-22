package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/marketplace/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/spacerouter/marketplace/docs"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/tea", controllers.GetTea)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("v1")
	{
		stack := controllers.StackController{
			DB: db,
		}
		v1.GET("stacks/", stack.GetAllStacks)

		v1.GET("stack/:id", stack.GetStackById)
		v1.GET("developer/:id", stack.GetStackByUserId)

		v1.GET("search/stack/*search", stack.GetStackSearch)

		v1.POST("/import", stack.ImportCompose)
		v1.POST("/fake_import", stack.FakeImport)
	}

	return router
}
