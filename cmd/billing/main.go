package main

import (
	"github.com/caarlos0/env/v11"
	"google.golang.org/grpc"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
	grpcTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/transport/grpc"
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

	err = db.AutoMigrate(&model.Bill{}, &model.Payment{})
	if err != nil {
		panic(err)
	}

	svc := billing.NewService(db)

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	server, err := grpcTransport.NewServer(cfg.GRPCAddress, grpcOpts, svc)
	if err != nil {
		return
	}

	server.Serve()
}
