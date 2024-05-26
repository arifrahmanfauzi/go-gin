package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/drops", controllers.GetDrops)
		api.GET("/aggregate", controllers.GetAggregateDrops)
	}
}
