package grpc

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

type Server struct {
	pb.UnimplementedOrderingServiceServer
	grpcServer *grpc.Server
	listener   net.Listener

	svc ordering.Service
}

func NewServer(
	address string, opts []grpc.ServerOption, service ordering.Service,
) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(opts...)

	srv := Server{
		grpcServer: grpcServer,
		listener:   listener,
		svc:        service,
	}

	pb.RegisterOrderingServiceServer(grpcServer, &srv)

	return &srv, nil
}

func (s *Server) Serve() {
	log.Printf("gRPC server listening on %s", s.listener.Addr())

	if err := s.grpcServer.Serve(s.listener); err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

func (s *Server) CreateFood(ctx context.Context, req *pb.CreateFoodRequest) (*pb.CreateFoodResponse, error) {
	input := req.Input

	food, err := s.svc.CreateFood(ctx, ordering.CreateFoodInput{
		Name:        input.Name,
		Description: input.Description,
		Price:       int(input.Price),
		ImageURL:    input.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	f, err := modelFoodTogRPCFood(food)
	if err != nil {
		return nil, err
	}

	return &pb.CreateFoodResponse{
		CreatedFood: f,
	}, nil
}

func (s *Server) GetFood(ctx context.Context, req *pb.GetFoodRequest) (*pb.GetFoodResponse, error) {
	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get food: %v", err)
	}

	food, err := s.svc.GetFood(ctx, uint(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get food: %v", err)
	}

	if food == nil {
		return nil, status.Errorf(codes.NotFound, "food not found")
	}

	f, err := modelFoodTogRPCFood(food)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get food: %v", err)
	}

	return &pb.GetFoodResponse{
		Food: f,
	}, nil
}

func (s *Server) GetFoods(ctx context.Context, _ *pb.GetFoodsRequest) (*pb.GetFoodsResponse, error) {
	foods, err := s.svc.GetFoods(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get foods: %v", err)
	}

	var fs []*pb.Food
	for _, food := range foods {
		f, err := modelFoodTogRPCFood(&food)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get foods: %v", err)
		}

		fs = append(fs, f)
	}

	return &pb.GetFoodsResponse{
		Foods: fs,
	}, nil
}

func (s *Server) UpdateFood(ctx context.Context, req *pb.UpdateFoodRequest) (*pb.UpdateFoodResponse, error) {
	input := req.Input

	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update food: %v", err)
	}

	var price int
	if input.Price == nil {
		price = -1
	} else {
		price = int(*input.Price)
	}

	food, err := s.svc.UpdateFood(ctx, uint(id), ordering.UpdateFoodInput{
		Name:        input.Name,
		Description: input.Description,
		Price:       &price,
		ImageURL:    input.ImageUrl,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update food: %v", err)
	}

	f, err := modelFoodTogRPCFood(food)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update food: %v", err)
	}

	return &pb.UpdateFoodResponse{
		UpdatedFood: f,
	}, nil
}

func (s *Server) DeleteFood(ctx context.Context, req *pb.DeleteFoodRequest) (*pb.DeleteFoodResponse, error) {
	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete food: %v", err)
	}

	err = s.svc.DeleteFood(ctx, uint(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete food: %v", err)
	}

	return &pb.DeleteFoodResponse{}, nil
}

func (s *Server) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	foodId, err := strconv.Atoi(req.FoodId.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart: %v", err)
	}

	cartItem, err := s.svc.AddToCart(ctx, req.UserId, uint(foodId), uint(req.Amount))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart: %v", err)
	}

	ci, err := modelCartItemTogRPCCartItem(cartItem)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart: %v", err)
	}

	return &pb.AddToCartResponse{
		AddedItem: ci,
	}, nil
}

func (s *Server) GetCartItems(ctx context.Context, req *pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error) {
	cartItems, err := s.svc.GetCartItems(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}

	var cItems []*pb.CartItem
	for _, cItem := range cartItems {
		ci, err := modelCartItemTogRPCCartItem(&cItem)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
		}

		cItems = append(cItems, ci)
	}

	return &pb.GetCartItemsResponse{
		Items: cItems,
	}, nil
}

func (s *Server) UpdateCartAmount(ctx context.Context, req *pb.UpdateCartAmountRequest) (*pb.UpdateCartAmountResponse, error) {
	cartItemId, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart: %v", err)
	}

	cartItem, err := s.svc.UpdateCartAmount(ctx, uint(cartItemId), uint(req.Amount))
	if err != nil {
		return nil, err
	}

	ci, err := modelCartItemTogRPCCartItem(cartItem)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart: %v", err)
	}

	return &pb.UpdateCartAmountResponse{
		UpdatedItem: ci,
	}, nil
}

func (s *Server) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	err := s.svc.ClearCart(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to clear cart: %v", err)
	}

	return &pb.ClearCartResponse{}, nil
}

func (s *Server) MakeOrder(ctx context.Context, req *pb.MakeOrderRequest) (*pb.MakeOrderResponse, error) {
	order, err := s.svc.MakeOrder(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
	}

	o, err := modelOrderTogRPCOrder(order)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
	}

	return &pb.MakeOrderResponse{
		CreatedOrder: o,
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
	}

	order, err := s.svc.GetOrder(ctx, uint(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
	}

	o, err := modelOrderTogRPCOrder(order)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
	}

	return &pb.GetOrderResponse{Order: o}, nil
}

func (s *Server) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	orders, err := s.svc.GetOrders(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders: %v", err)
	}

	var os []*pb.Order
	for _, order := range orders {
		o, err := modelOrderTogRPCOrder(&order)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to make order: %v", err)
		}
		os = append(os, o)
	}

	return &pb.GetOrdersResponse{Orders: os}, nil
}
