package grpc

import (
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing"
)

type Server struct {
	pb.UnimplementedBillingServiceServer
	grpcServer *grpc.Server
	listener   net.Listener

	svc billing.Service
}

func NewServer(
	address string, opts []grpc.ServerOption, service billing.Service,
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

	pb.RegisterBillingServiceServer(grpcServer, &srv)

	return &srv, nil
}

func (s *Server) Serve() {
	if err := s.grpcServer.Serve(s.listener); err != nil {
		panic(err)
	}
}

func (s *Server) CreateBill(
	ctx context.Context, req *pb.CreateBillRequest,
) (*pb.CreateBillResponse, error) {
	input := req.Input

	bill, err := s.svc.CreateBill(ctx, input.UserId, input.OrderId, int(input.Amount))
	if err != nil {
		return nil, err
	}

	b := pb.Bill{
		Id:        strconv.Itoa(int(bill.ID)),
		CreatedAt: timestamppb.New(bill.CreatedAt),
		UpdatedAt: timestamppb.New(bill.UpdatedAt),

		UserId:  bill.UserId,
		OrderId: bill.OrderId,
		Amount:  int32(bill.Amount),
	}

	return &pb.CreateBillResponse{
		CreatedBill: &b,
	}, nil
}
