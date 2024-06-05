package handler

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const billTarget = "localhost:50051"
const notificationTarget = "localhost:50052"
const orderTarget = "locahost:50053"
const ratingTarget = "localhost:50054"
const userTarget = "localhost:50055"

func newClient(target string) (*grpc.ClientConn, error) {
	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(target, grpcOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
