package ordering

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
)

func (svc *service) AddToCart(ctx context.Context, userId string, foodId uint, amount uint) (*model.CartItem, error) {
	item := model.NewCartItem(userId, foodId, amount)

	result := svc.db.Clauses(clause.Returning{}).Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (svc *service) GetCartItems(ctx context.Context, userId string) ([]model.CartItem, error) {
	var items []model.CartItem

	result := svc.db.Where("user_id = ?", userId).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (svc *service) UpdateCartAmount(ctx context.Context, cartItemId uint, amount uint) (*model.CartItem, error) {
	var item *model.CartItem

	svc.db.First(&item, cartItemId)

	if item == nil {
		return nil, gorm.ErrRecordNotFound
	}

	item.Amount = amount
	result := svc.db.Clauses(clause.Returning{}).Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (svc *service) ClearCart(ctx context.Context, userId string) error {
	err := svc.db.Delete("user_id = ?", userId).Error

	return err
}
