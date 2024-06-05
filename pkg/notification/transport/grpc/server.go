package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/notification"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	grpcServer *grpc.Server
	listener   net.Listener

	svc notification.Service
}

func NewServer(
	address string, opts []grpc.ServerOption, service notification.Service,
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

	pb.RegisterNotificationServiceServer(grpcServer, &srv)

	return &srv, nil
}

func (s *Server) Serve() {
	if err := s.grpcServer.Serve(s.listener); err != nil {
		panic(err)
	}
}

func (s *Server) SendPayPaymentNotification(ctx context.Context, request *pb.SendPayPaymentNotificationRequest) (*pb.SendPayPaymentNotificationResponse, error) {
	err := s.svc.SendPayPaymentNotification(ctx, request.UserId, int(request.Amount))
	if err != nil {
		return nil, err
	}

	return &pb.SendPayPaymentNotificationResponse{}, nil
}
