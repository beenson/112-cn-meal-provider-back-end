package main

import (
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
	grpcTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/transport/grpc"
)

func main() {

	db, err := gorm.Open(sqlite.Open("billing.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Bill{}, &model.Payment{})
	if err != nil {
		panic(err)
	}

	svc := billing.NewService(db)

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	server, err := grpcTransport.NewServer("0.0.0.0:50051", grpcOpts, svc)
	if err != nil {
		return
	}

	server.Serve()
}
