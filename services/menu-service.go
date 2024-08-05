package services

import (
	"github.com/nomad483/restaurant-managment/models"
	"github.com/nomad483/restaurant-managment/repositories"
)

var menuRepository repositories.MenuRepository = repositories.NewMenuRepository()

type MenuService interface {
	GetMenus() (models.ResponseMenuList, error)
	GetMenuById(id string) (models.ResponseMenu, error)
	CreateMenu(menu models.Menu) (models.ResponseMenu, error)
	UpdateMenu(id string, menu models.UpdateMenu) (models.ResponseMenu, error)
	DeleteMenu(id string) error

	mapMenuToResponseMenu(menu models.Menu) models.ResponseMenu
}

type menuService struct{}

func NewMenuService() MenuService {
	return &menuService{}
}

func (service *menuService) mapMenuToResponseMenu(menu models.Menu) models.ResponseMenu {
	return models.ResponseMenu{
		ID:        menu.ID,
		Name:      menu.Name,
		Category:  menu.Category,
		StartDate: menu.StartDate,
		EndDate:   menu.EndDate,
	}
}

func (service *menuService) GetMenus() (models.ResponseMenuList, error) {
	menus, err := menuRepository.GetMenus()
	if err != nil {
		return models.ResponseMenuList{}, err
	}

	var responseMenus []models.ResponseMenu

	for _, menu := range menus {
		responseMenus = append(responseMenus, service.mapMenuToResponseMenu(menu))
	}

	return models.ResponseMenuList{Data: responseMenus}, nil
}

func (service *menuService) GetMenuById(id string) (models.ResponseMenu, error) {
	menu, err := menuRepository.GetMenuBuId(id)
	if err != nil {
		return models.ResponseMenu{}, err
	}

	return service.mapMenuToResponseMenu(menu), nil
}

func (service *menuService) CreateMenu(menu models.Menu) (models.ResponseMenu, error) {
	newMenu, err := menuRepository.CreateMenu(menu)
	if err != nil {
		return models.ResponseMenu{}, err
	}

	return service.mapMenuToResponseMenu(newMenu), nil
}

func (service *menuService) UpdateMenu(id string, menu models.UpdateMenu) (models.ResponseMenu, error) {
	updatedMenu, err := menuRepository.UpdateMenu(id, menu)
	if err != nil {
		return models.ResponseMenu{}, err
	}

	return service.mapMenuToResponseMenu(updatedMenu), nil
}

func (service *menuService) DeleteMenu(id string) error {
	return menuRepository.DeleteMenu(id)
}
