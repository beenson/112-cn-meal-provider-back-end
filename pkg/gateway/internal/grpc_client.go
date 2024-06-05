package internal

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(target string) (*grpc.ClientConn, error) {
	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(target, grpcOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
