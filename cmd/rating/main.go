package main

import (
	"fmt"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/rating/manager"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/rating/model"

	"google.golang.org/grpc"
)

func main() {
	db, err := model.CreateConnection()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(model.GetModels()...)
	fmt.Println("Migrated successfully")

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	m := manager.NewManager("127.0.0.1:50050", grpcOpts, db)
	go m.Serve()

	channel := make(chan string)
	<-channel
	fmt.Println("serve called")
}
