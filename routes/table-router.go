package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	tableGroup := incomingRoutes.Group("/api/tables")

	tableGroup.GET("/", controllers.GetTables())
	tableGroup.GET("/:id", controllers.GetTable())
	tableGroup.POST("/", controllers.CreateTable())
	tableGroup.PATCH("/:id", controllers.UpdateTable())
	tableGroup.DELETE("/:id", controllers.DeleteTable())
}
