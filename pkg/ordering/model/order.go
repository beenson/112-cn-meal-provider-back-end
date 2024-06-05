package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserId string
	Items  []OrderItem `gorm:"foreignKey:OrderId"`

	SubTotal int
}

func NewOrder(userId string) Order {
	return Order{
		UserId:   userId,
		SubTotal: 0,
	}
}

func (order Order) AddItem(item OrderItem) {
	order.Items = append(order.Items, item)
}

func (order Order) AddItemFromCart(cartItems []CartItem) {
	subTotal := 0

	for _, item := range cartItems {
		order.AddItem(OrderItem{
			FoodId: item.FoodId,
			Amount: item.Amount,
		})

		subTotal += item.Food.Price * int(item.Amount)
	}

	order.SubTotal += subTotal
}

type OrderItem struct {
	gorm.Model

	OrderId uint
	FoodId  uint
	Amount  uint
}
