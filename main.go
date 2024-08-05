package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	routes.UserRoutes(router)

	//router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
