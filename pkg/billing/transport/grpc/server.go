package grpc

import (
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	b, err := modelBillTogRPCBill(bill)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBillResponse{
		CreatedBill: b,
	}, nil
}

func (s *Server) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	input := req.Input

	payment, err := s.svc.CreatePayment(ctx, input.UserId, int(input.Amount))
	if err != nil {
		return nil, err
	}

	p, err := modelPaymentTogRPCPayment(payment)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePaymentResponse{
		CreatedPayment: p,
	}, nil
}

func (s *Server) GetBill(ctx context.Context, req *pb.GetBillRequest) (*pb.GetBillResponse, error) {
	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bill, err := s.svc.GetBill(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	b, err := modelBillTogRPCBill(bill)
	if err != nil {
		return nil, err
	}

	return &pb.GetBillResponse{
		Bill: b,
	}, nil
}

func (s *Server) GetBills(ctx context.Context, req *pb.GetBillsRequest) (*pb.GetBillsResponse, error) {
	params := req.QueryParams

	modelBills, err := s.svc.GetBills(ctx, *params.UserId)
	if err != nil {
		return nil, err
	}

	pbBills := make([]*pb.Bill, len(modelBills))
	for i, bill := range modelBills {
		pbBills[i], err = modelBillTogRPCBill(&bill)
	}

	return &pb.GetBillsResponse{
		Bills: pbBills,
	}, nil
}

func (s *Server) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	id, err := strconv.Atoi(req.Id.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	payment, err := s.svc.GetPayment(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	p, err := modelPaymentTogRPCPayment(payment)
	if err != nil {
		return nil, err
	}

	return &pb.GetPaymentResponse{
		Payment: p,
	}, nil
}

func (s *Server) GetPayments(ctx context.Context, req *pb.GetPaymentsRequest) (*pb.GetPaymentsResponse, error) {
	params := req.QueryParams

	modelPayments, err := s.svc.GetPayments(ctx, *params.UserId)
	if err != nil {
		return nil, err
	}

	pbPayments := make([]*pb.Payment, len(modelPayments))
	for i, p := range modelPayments {
		pbPayments[i], err = modelPaymentTogRPCPayment(&p)
	}

	return &pb.GetPaymentsResponse{
		Payments: pbPayments,
	}, nil
}
