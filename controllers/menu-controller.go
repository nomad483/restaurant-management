package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/services"
	"net/http"
)

var menuService services.MenuService = services.NewMenuService()

type MenuController interface {
	GetMenus() gin.HandlerFunc
	GetMenu() gin.HandlerFunc
	CreateMenu() gin.HandlerFunc
	UpdateMenu() gin.HandlerFunc
	DeleteMenu() gin.HandlerFunc
}

type menuController struct{}

func NewMenuController() MenuController {
	return &menuController{}
}

func (controller *menuController) GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseMenus, err := menuService.GetMenus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responseMenus)
	}
}

func (controller *menuController) GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		responseMenu, err := menuService.GetMenuById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responseMenu)
	}
}

func (controller *menuController) CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		responseMenus, err := menuService.CreateMenu(menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responseMenus)
	}
}

func (controller *menuController) UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var menu models.UpdateMenu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedMenu, err := menuService.UpdateMenu(id, menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedMenu)
	}
}

func (controller *menuController) DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := menuService.DeleteMenu(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Menu deleted"})
	}
}
