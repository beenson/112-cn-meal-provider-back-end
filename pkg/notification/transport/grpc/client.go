package grpc

import (
	"context"

	"google.golang.org/grpc"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/notification"
)

type gRPCClient struct {
	gRPC pb.NotificationServiceClient
}

func NewClient(target string, opts ...grpc.DialOption) (notification.Service, error) {
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	gRPC := pb.NewNotificationServiceClient(conn)

	return &gRPCClient{
		gRPC: gRPC,
	}, nil
}

func (g *gRPCClient) SendPayPaymentNotification(ctx context.Context, userId string, amountToPay int) error {
	request := &pb.SendPayPaymentNotificationRequest{
		UserId: userId,
		Amount: int32(amountToPay),
	}

	_, err := g.gRPC.SendPayPaymentNotification(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
