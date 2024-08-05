package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

var menuController controllers.MenuController = controllers.NewMenuController()

func MenuRoutes(incomingRoutes *gin.Engine) {
	menuGroup := incomingRoutes.Group("/api/menu")

	menuGroup.GET("/", menuController.GetMenus())
	menuGroup.GET("/:id", menuController.GetMenu())
	menuGroup.POST("/", menuController.CreateMenu())
	menuGroup.PATCH("/:id", menuController.UpdateMenu())
	menuGroup.DELETE("/:id", menuController.DeleteMenu())
}
