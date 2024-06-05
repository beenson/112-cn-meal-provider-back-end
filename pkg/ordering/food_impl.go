package ordering

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
)

func (svc *service) CreateFood(ctx context.Context, input CreateFoodInput) (*model.Food, error) {
	f := model.NewFood(input.Name, input.Description, input.Price, input.ImageURL)

	result := svc.db.Clauses(clause.Returning{}).Create(&f)
	if result.Error != nil {
		return nil, result.Error
	}

	return &f, nil
}

func (svc *service) GetFood(ctx context.Context, id uint) (*model.Food, error) {
	f := &model.Food{}

	result := svc.db.First(f, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return f, nil
}

func (svc *service) GetFoods(ctx context.Context) ([]model.Food, error) {
	var foods []model.Food

	result := svc.db.Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}

	return foods, nil
}

func (svc *service) UpdateFood(ctx context.Context, id uint, input UpdateFoodInput) (*model.Food, error) {
	var f *model.Food
	svc.db.First(&f, id)
	if f == nil {
		return nil, gorm.ErrRecordNotFound
	}

	if input.Name != nil {
		f.Name = *input.Name
	}

	if input.Description != nil {
		f.Description = *input.Description
	}

	if input.Price != nil {
		f.Price = *input.Price
	}

	if input.ImageURL != nil {
		f.ImageURL = *input.ImageURL
	}

	err := svc.db.Clauses(clause.Returning{}).Save(&f).Error
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (svc *service) DeleteFood(ctx context.Context, id uint) error {
	result := svc.db.Delete(&model.Food{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
