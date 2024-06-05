package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/rating/manager"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/rating/model"

	"google.golang.org/grpc"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	db, err := cfg.DB.GetGormDB()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(model.GetModels()...)
	fmt.Println("Migrated successfully")

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	m := manager.NewManager(cfg.GRPCAddress, grpcOpts, db)
	go m.Serve()

	channel := make(chan string)
	<-channel
	fmt.Println("serve called")
}
