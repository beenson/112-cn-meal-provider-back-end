package grpc

import (
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func gRPCOrderToModelOrder(order *pb.Order) (*model.Order, error) {
	if order == nil {
		return nil, nil
	}

	id, err := strconv.Atoi(order.Id.Id)
	if err != nil {
		return nil, err
	}

	items := make([]model.OrderItem, len(order.Items))
	for i, item := range order.Items {
		oi, err := gRPCOrderItemToModelOrderItem(item)
		if err != nil {
			return nil, err
		}

		items[i] = *oi
	}

	return &model.Order{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: order.CreatedAt.AsTime(),
			UpdatedAt: order.UpdatedAt.AsTime(),
		},
		UserId:   order.UserId,
		SubTotal: int(order.Subtotal),

		Items: items,
	}, nil
}

func modelOrderTogRPCOrder(order *model.Order) (*pb.Order, error) {
	if order == nil {
		return nil, nil
	}

	items := make([]*pb.OrderItem, len(order.Items))
	for i, item := range order.Items {
		oi, err := modelOrderItemTogRPCOrderItem(&item)
		if err != nil {
			return nil, err
		}

		items[i] = oi
	}

	return &pb.Order{
		Id:        &pb.OrderID{Id: strconv.Itoa(int(order.ID))},
		CreatedAt: timestamppb.New(order.CreatedAt),
		UpdatedAt: timestamppb.New(order.UpdatedAt),

		UserId:   order.UserId,
		Subtotal: int32(order.SubTotal),

		Items: items,
	}, nil
}

func gRPCOrderItemToModelOrderItem(orderItem *pb.OrderItem) (*model.OrderItem, error) {
	if orderItem == nil {
		return nil, nil
	}

	id, err := strconv.Atoi(orderItem.Id.Id)
	if err != nil {
		return nil, err
	}

	orderId, err := strconv.Atoi(orderItem.OrderId)
	if err != nil {
		return nil, err
	}

	foodId, err := strconv.Atoi(orderItem.FoodId)
	if err != nil {
		return nil, err
	}

	return &model.OrderItem{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: orderItem.CreatedAt.AsTime(),
			UpdatedAt: orderItem.UpdatedAt.AsTime(),
		},
		OrderId: uint(orderId),
		FoodId:  uint(foodId),
		Amount:  uint(orderItem.Amount),
	}, nil
}

func modelOrderItemTogRPCOrderItem(orderItem *model.OrderItem) (*pb.OrderItem, error) {
	if orderItem == nil {
		return nil, nil
	}

	return &pb.OrderItem{
		Id:        &pb.OrderItemID{Id: strconv.Itoa(int(orderItem.ID))},
		CreatedAt: timestamppb.New(orderItem.CreatedAt),
		UpdatedAt: timestamppb.New(orderItem.UpdatedAt),

		OrderId: strconv.Itoa(int(orderItem.OrderId)),
		FoodId:  strconv.Itoa(int(orderItem.FoodId)),
		Amount:  uint32(orderItem.Amount),
	}, nil
}
