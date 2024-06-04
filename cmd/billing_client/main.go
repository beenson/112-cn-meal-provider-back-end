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

	userId := "1234"

	bill, err := clientSvc.CreateBill(context.Background(), userId, "abc", 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create Bill: %+v\n", bill)

	{
		b, err := clientSvc.GetBill(context.Background(), bill.ID)
		if err != nil {
			return
		}

		fmt.Printf("Get Bill: %+v\n", b)
	}

	{
		_, err := clientSvc.CreateBill(context.Background(), userId, "def", 20)
		if err != nil {
			panic(err)
		}

		b, err := clientSvc.GetBills(context.Background(), userId)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Get Bills: %+v\n", b)
	}
}
