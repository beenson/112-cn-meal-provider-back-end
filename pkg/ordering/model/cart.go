package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	UserId string

	FoodId uint
	Food   Food `gorm:"foreignKey:FoodId"`

	Amount uint
}

func NewCartItem(userId string, foodId uint, amount uint) CartItem {
	return CartItem{
		UserId: userId,
		FoodId: foodId,
		Amount: amount,
	}
}
