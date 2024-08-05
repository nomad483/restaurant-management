package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userGroup := incomingRoutes.Group("/api/users")

	userGroup.GET("/", controllers.GetUsers())
	userGroup.GET("/:id", controllers.GetUser())
	userGroup.POST("/signup", controllers.SignUp())
	userGroup.POST("/login", controllers.SignIn())
}
