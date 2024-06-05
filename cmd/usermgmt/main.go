package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/manager"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/model"

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

	db.Create(&model.User{
		Uid:        "87877",
		Group:      "admin",
		Username:   "username",
		Department: "cs",
		Email:      "a@b.c",
		Password:   "password",
	})

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	m := manager.NewManager(cfg.GRPCAddress, grpcOpts, db)
	go m.Serve()

	channel := make(chan string)
	<-channel
	fmt.Println("serve called")
}
