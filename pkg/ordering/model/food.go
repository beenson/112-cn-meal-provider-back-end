package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model

	Name        string
	Description string
	Price       int
	ImageURL    string
}

func NewFood(name string, description string, price int, imageURL string) Food {
	return Food{
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL,
	}
}
