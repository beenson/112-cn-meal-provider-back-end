package grpc

import (
	"context"
	"gorm.io/gorm"
	"strconv"

	"google.golang.org/grpc"

	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
)

type gRPCClient struct {
	gRPC pb.BillingServiceClient
}

func NewClient(target string, opts ...grpc.DialOption) (billing.Service, error) {
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	gRPC := pb.NewBillingServiceClient(conn)

	return &gRPCClient{
		gRPC: gRPC,
	}, nil
}

func (c *gRPCClient) CreateBill(ctx context.Context, userId string, orderId string, amount int) (*model.Bill, error) {
	request := &pb.CreateBillRequest{
		Input: &pb.CreateBillInput{
			UserId:  userId,
			OrderId: orderId,
			Amount:  int32(amount),
		},
	}

	resp, err := c.gRPC.CreateBill(ctx, request)
	if err != nil {
		return nil, err
	}

	b, err := gRPCBillToModelBill(resp.CreatedBill)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *gRPCClient) CreatePayment(ctx context.Context, userId string, amount int) (*model.Payment, error) {
	panic("implement me")
}

func (c *gRPCClient) GetBill(ctx context.Context, id uint) (*model.Bill, error) {
	panic("implement me")
}

func (c *gRPCClient) GetBills(ctx context.Context, userId string) ([]model.Bill, error) {
	panic("implement me")
}

func (c *gRPCClient) GetPayment(ctx context.Context, id uint) (*model.Payment, error) {
	panic("implement me")
}

func (c *gRPCClient) GetPayments(ctx context.Context, userId string) ([]model.Payment, error) {
	panic("implement me")
}

func gRPCBillToModelBill(bill *pb.Bill) (*model.Bill, error) {
	id, err := strconv.Atoi(bill.Id)
	if err != nil {
		return nil, err
	}

	return &model.Bill{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: bill.CreatedAt.AsTime(),
			UpdatedAt: bill.UpdatedAt.AsTime(),
		},
		UserId:  bill.UserId,
		OrderId: bill.OrderId,
		Amount:  int(bill.Amount),
	}, nil
}
