package grpc

import (
	"context"
	"google.golang.org/grpc"
	"strconv"

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
	request := &pb.CreatePaymentRequest{
		Input: &pb.CreatePaymentInput{
			UserId: userId,
			Amount: int32(amount),
		},
	}

	resp, err := c.gRPC.CreatePayment(ctx, request)
	if err != nil {
		return nil, err
	}

	p, err := gRPCPaymentToModelPayment(resp.CreatedPayment)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *gRPCClient) GetBill(ctx context.Context, id uint) (*model.Bill, error) {
	request := &pb.GetBillRequest{
		Id: &pb.BillId{Id: strconv.Itoa(int(id))},
	}

	resp, err := c.gRPC.GetBill(ctx, request)
	if err != nil {
		return nil, err
	}

	b, err := gRPCBillToModelBill(resp.Bill)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *gRPCClient) GetBills(ctx context.Context, userId string) ([]model.Bill, error) {
	request := &pb.GetBillsRequest{
		QueryParams: &pb.QueryBillParams{
			UserId: &userId,
		},
	}

	resp, err := c.gRPC.GetBills(ctx, request)
	if err != nil {
		return nil, err
	}

	bills := make([]model.Bill, len(resp.Bills))
	for i, bill := range resp.Bills {
		b, err := gRPCBillToModelBill(bill)
		if err != nil {
			return nil, err
		}

		bills[i] = *b
	}

	return bills, nil
}

func (c *gRPCClient) GetPayment(ctx context.Context, id uint) (*model.Payment, error) {
	panic("implement me")
}

func (c *gRPCClient) GetPayments(ctx context.Context, userId string) ([]model.Payment, error) {
	panic("implement me")
}
