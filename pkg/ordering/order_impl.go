package ordering

import (
	"context"
	"gorm.io/gorm/clause"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
)

func (svc *service) MakeOrder(ctx context.Context, userId string) (*model.Order, error) {
	cartItems, err := svc.GetCartItems(ctx, userId)
	if err != nil {
		return nil, err
	}

	order := model.NewOrder(userId)
	order.AddItemFromCart(cartItems)

	result := svc.db.Clauses(clause.Returning{}).Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (svc *service) GetOrder(ctx context.Context, orderId uint) (*model.Order, error) {
	var order *model.Order

	result := svc.db.First(&order, orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (svc *service) GetOrders(ctx context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order

	result := svc.db.Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
