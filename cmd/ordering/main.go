package main

import (
	"github.com/caarlos0/env/v11"
	"google.golang.org/grpc"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
	grpcTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/transport/grpc"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	db, err := cfg.DB.GetGormDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Food{}, &model.CartItem{}, &model.OrderItem{}, &model.Order{})
	if err != nil {
		panic(err)
	}

	svc := ordering.NewService(db)

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	server, err := grpcTransport.NewServer(cfg.GRPCAddress, grpcOpts, svc)
	if err != nil {
		return
	}

	go server.Serve()

	internal.WaitUntilShutdownSignal()
	server.Stop()
}
