package ordering

import (
	"context"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
)

type CreateFoodInput struct {
	Name        string
	Description string
	Price       int
	ImageURL    string
}

type UpdateFoodInput struct {
	Name        *string
	Description *string
	Price       *int
	ImageURL    *string
}

type Service interface {
	CreateFood(ctx context.Context, input CreateFoodInput) (*model.Food, error)

	GetFood(ctx context.Context, id uint) (*model.Food, error)
	GetFoods(ctx context.Context) ([]model.Food, error)

	UpdateFood(ctx context.Context, id uint, input UpdateFoodInput) (*model.Food, error)
	DeleteFood(ctx context.Context, id uint) error

	AddToCart(ctx context.Context, userId string, foodId uint, amount uint) (*model.CartItem, error)
	GetCartItems(ctx context.Context, userId string) ([]model.CartItem, error)
	UpdateCartAmount(ctx context.Context, cartItemId uint, amount uint) (*model.CartItem, error)
	ClearCart(ctx context.Context, userId string) error

	MakeOrder(ctx context.Context, userId string) (*model.Order, error)

	GetOrder(ctx context.Context, orderId uint) (*model.Order, error)
	GetOrders(ctx context.Context, userId string) ([]model.Order, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
