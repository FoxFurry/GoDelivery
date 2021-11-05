package dto

import "github.com/foxfurry/go_delivery/internal/domain/entity"

type Restaurant struct {
	Name string `json:"name"`
	MenuItems int `json:"menu_items"`
	RestaurantMenu []entity.Food `json:"menu"`
}
