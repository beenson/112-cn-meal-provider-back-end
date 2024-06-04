package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gRPCTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/transport/grpc"
)

func main() {
	clientSvc, err := gRPCTransport.NewClient(
		"localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	bill, err := clientSvc.CreateBill(context.Background(), "1234", "abc", 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create Bill: %+v\n", bill)
}
