package manager

import (
	"fmt"
	"net"

	protocol "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/user_mgmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GRPCManager struct {
	protocol.UnimplementedUserManagementServiceServer
	grpcServer *grpc.Server
	listener	net.Listener
	db         *gorm.DB
}

func NewManager(address string, opts []grpc.ServerOption, db *gorm.DB) *GRPCManager {

	if lis, err := net.Listen("tcp", address); err != nil {
		fmt.Printf("net.Listen() failed: %v", err)
		panic(err)
	}
	svc := &GRPCManger{
		listener:   lis,
		grpcServer: grpc.NewServer(opts...),
		db:         db,
	}
	protocol.RegisterUserManagementServiceServer(svc.grpcServer, svc)

	return svc
}

func (s *GRPCManger) Serve() {
	if err := s.grpcServer.Serve(s.lis); err != nil {
		panic(err)
	}
}