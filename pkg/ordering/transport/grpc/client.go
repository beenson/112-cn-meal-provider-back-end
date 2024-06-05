package grpc

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

type gRPCClient struct {
	gRPC pb.OrderingServiceClient
}

func NewClient(target string, opts ...grpc.DialOption) (ordering.Service, error) {
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	gRPC := pb.NewOrderingServiceClient(conn)

	return &gRPCClient{
		gRPC: gRPC,
	}, nil
}

func (g *gRPCClient) CreateFood(ctx context.Context, input ordering.CreateFoodInput) (*model.Food, error) {
	req := &pb.CreateFoodRequest{
		Input: &pb.CreateFoodInput{
			Name:        input.Name,
			Description: input.Description,
			Price:       int32(input.Price),
			ImageUrl:    input.ImageURL,
		},
	}

	resp, err := g.gRPC.CreateFood(ctx, req)
	if err != nil {
		return nil, err
	}

	food, err := gRPCFoodToModelFood(resp.CreatedFood)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (g *gRPCClient) GetFood(ctx context.Context, id uint) (*model.Food, error) {
	req := &pb.GetFoodRequest{
		Id: &pb.FoodID{
			Id: strconv.Itoa(int(id)),
		},
	}

	resp, err := g.gRPC.GetFood(ctx, req)
	if err != nil {
		return nil, err
	}

	food, err := gRPCFoodToModelFood(resp.Food)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (g *gRPCClient) GetFoods(ctx context.Context) ([]model.Food, error) {
	req := &pb.GetFoodsRequest{}

	resp, err := g.gRPC.GetFoods(ctx, req)
	if err != nil {
		return nil, err
	}

	var foods []model.Food
	for _, f := range resp.Foods {
		food, err := gRPCFoodToModelFood(f)
		if err != nil {
			return nil, err
		}

		foods = append(foods, *food)
	}

	return foods, nil
}

func (g *gRPCClient) UpdateFood(ctx context.Context, id uint, input ordering.UpdateFoodInput) (*model.Food, error) {
	var price *int32

	if input.Price != nil {
		price = new(int32)
		*price = int32(*input.Price)
	}

	req := &pb.UpdateFoodRequest{
		Id: &pb.FoodID{
			Id: strconv.Itoa(int(id)),
		},
		Input: &pb.UpdateFoodInput{
			Name:        input.Name,
			Description: input.Description,
			Price:       price,
			ImageUrl:    input.ImageURL,
		},
	}

	resp, err := g.gRPC.UpdateFood(ctx, req)
	if err != nil {
		return nil, err
	}

	food, err := gRPCFoodToModelFood(resp.UpdatedFood)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (g *gRPCClient) DeleteFood(ctx context.Context, id uint) error {
	req := &pb.DeleteFoodRequest{
		Id: &pb.FoodID{
			Id: strconv.Itoa(int(id)),
		},
	}

	_, err := g.gRPC.DeleteFood(ctx, req)
	return err
}

func (g *gRPCClient) AddToCart(ctx context.Context, userId string, foodId uint, amount uint) (*model.CartItem, error) {
	req := &pb.AddToCartRequest{
		UserId: userId,
		FoodId: &pb.FoodID{
			Id: strconv.Itoa(int(foodId)),
		},
		Amount: uint32(amount),
	}

	resp, err := g.gRPC.AddToCart(ctx, req)
	if err != nil {
		return nil, err
	}

	cartItem, err := gRPCCartItemToModelCartItem(resp.AddedItem)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (g *gRPCClient) GetCartItems(ctx context.Context, userId string) ([]model.CartItem, error) {
	req := &pb.GetCartItemsRequest{
		UserId: userId,
	}

	resp, err := g.gRPC.GetCartItems(ctx, req)
	if err != nil {
		return nil, err
	}

	var cartItems []model.CartItem
	for _, i := range resp.Items {
		cartItem, err := gRPCCartItemToModelCartItem(i)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, *cartItem)
	}

	return cartItems, nil
}

func (g *gRPCClient) UpdateCartAmount(ctx context.Context, cartItemId uint, amount uint) (*model.CartItem, error) {
	req := &pb.UpdateCartAmountRequest{
		Id: &pb.CartItemID{
			Id: strconv.Itoa(int(cartItemId)),
		},
		Amount: uint32(amount),
	}

	resp, err := g.gRPC.UpdateCartAmount(ctx, req)
	if err != nil {
		return nil, err
	}

	cartItem, err := gRPCCartItemToModelCartItem(resp.UpdatedItem)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (g *gRPCClient) ClearCart(ctx context.Context, userId string) error {
	req := &pb.ClearCartRequest{
		UserId: userId,
	}

	_, err := g.gRPC.ClearCart(ctx, req)
	return err
}

func (g *gRPCClient) MakeOrder(ctx context.Context, userId string) (*model.Order, error) {
	req := &pb.MakeOrderRequest{
		UserId: userId,
	}

	resp, err := g.gRPC.MakeOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	order, err := gRPCOrderToModelOrder(resp.CreatedOrder)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (g *gRPCClient) GetOrder(ctx context.Context, orderId uint) (*model.Order, error) {
	req := &pb.GetOrderRequest{
		Id: &pb.OrderID{
			Id: strconv.Itoa(int(orderId)),
		},
	}

	resp, err := g.gRPC.GetOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	order, err := gRPCOrderToModelOrder(resp.Order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (g *gRPCClient) GetOrders(ctx context.Context, userId string) ([]model.Order, error) {
	req := &pb.GetOrdersRequest{
		UserId: userId,
	}

	resp, err := g.gRPC.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}

	var orders []model.Order
	for _, o := range resp.Orders {
		order, err := gRPCOrderToModelOrder(o)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return orders, nil
}
