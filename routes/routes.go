package routes

import (
	"github.com/gin-gonic/gin"
	"go-starter/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/drops", controllers.GetDrops)
	}
}
