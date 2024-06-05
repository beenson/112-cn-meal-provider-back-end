package model

import "gorm.io/gorm"

type Bill struct {
	gorm.Model

	UserId  string
	OrderId string
	Amount  int
}

func NewBill(userId string, orderId string, amount int) Bill {
	return Bill{
		UserId:  userId,
		OrderId: orderId,
		Amount:  amount,
	}
}
