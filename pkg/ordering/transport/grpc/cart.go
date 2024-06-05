package grpc

import (
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func gRPCCartItemToModelCartItem(cartItem *pb.CartItem) (*model.CartItem, error) {
	if cartItem == nil {
		return nil, nil
	}

	id, err := strconv.Atoi(cartItem.Id.Id)
	if err != nil {
		return nil, err
	}

	foodId, err := strconv.Atoi(cartItem.FoodId)
	if err != nil {
		return nil, err
	}

	food, err := gRPCFoodToModelFood(cartItem.Food)
	if err != nil {
		return nil, err
	}

	if food == nil {
		food = &model.Food{}
	}

	return &model.CartItem{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: cartItem.CreatedAt.AsTime(),
			UpdatedAt: cartItem.UpdatedAt.AsTime(),
		},
		FoodId: uint(foodId),
		Food:   *food,
		Amount: uint(cartItem.Amount),
	}, nil
}

func modelCartItemTogRPCCartItem(cartItem *model.CartItem) (*pb.CartItem, error) {
	if cartItem == nil {
		return nil, nil
	}

	var food *pb.Food = nil
	if cartItem.Food.ID != 0 {
		var err error
		food, err = modelFoodTogRPCFood(&cartItem.Food)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CartItem{
		Id:        &pb.CartItemID{Id: strconv.Itoa(int(cartItem.ID))},
		CreatedAt: timestamppb.New(cartItem.CreatedAt),
		UpdatedAt: timestamppb.New(cartItem.UpdatedAt),

		FoodId: strconv.Itoa(int(cartItem.FoodId)),
		Food:   food,
		Amount: uint32(cartItem.Amount),
	}, nil
}
