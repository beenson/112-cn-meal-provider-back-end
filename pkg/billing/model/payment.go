package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model

	UserId string
	Amount int
}

func NewPayment(userId string, amount int) Payment {
	return Payment{
		UserId: userId,
		Amount: amount,
	}
}
