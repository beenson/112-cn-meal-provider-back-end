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

	p, err := clientSvc.GetPayment(context.Background(), 8888)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get Payment with 8888: %+v\n", p)

	payments, err := clientSvc.GetPayments(context.Background(), "8888")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get Payments with 8888: %+v\n", payments)
}
