package manager

import (
	"fmt"
	"net"

	protocol "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/rating"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GRPCManager struct {
	protocol.UnimplementedFoodRatingServiceServer
	grpcServer *grpc.Server
	listener   net.Listener
	db         *gorm.DB
}

func NewManager(address string, opts []grpc.ServerOption, db *gorm.DB) *GRPCManager {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("net.Listen() failed: %v", err)
		panic(err)
	}
	svc := &GRPCManager{
		listener:   lis,
		grpcServer: grpc.NewServer(opts...),
		db:         db,
	}
	protocol.RegisterFoodRatingServiceServer(svc.grpcServer, svc)

	return svc
}

func (s *GRPCManager) Serve() {
	if err := s.grpcServer.Serve(s.listener); err != nil {
		panic(err)
	}
}
