package grpc

import (
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func gRPCFoodToModelFood(food *pb.Food) (*model.Food, error) {
	if food == nil {
		return nil, nil
	}

	id, err := strconv.Atoi(food.Id.Id)
	if err != nil {
		return nil, err
	}

	return &model.Food{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: food.CreatedAt.AsTime(),
			UpdatedAt: food.UpdatedAt.AsTime(),
		},
		Name:        food.Name,
		Description: food.Description,
		Price:       int(food.Price),
		ImageURL:    food.ImageUrl,
	}, nil
}

func modelFoodTogRPCFood(food *model.Food) (*pb.Food, error) {
	if food == nil {
		return nil, nil
	}

	return &pb.Food{
		Id:        &pb.FoodID{Id: strconv.Itoa(int(food.ID))},
		CreatedAt: timestamppb.New(food.CreatedAt),
		UpdatedAt: timestamppb.New(food.UpdatedAt),

		Name:        food.Name,
		Description: food.Description,
		Price:       int32(food.Price),
		ImageUrl:    food.ImageURL,
	}, nil
}
