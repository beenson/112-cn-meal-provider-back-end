package main

import (
	"github.com/caarlos0/env/v11"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/sender/email"
	grpcTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/transport/grpc"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/userinfo"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	userProvider, err := userinfo.NewProvider(
		cfg.Service.UserMgmtTarget, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	mailSender := email.NewFakeEmailSender()

	svc := notification.NewService(userProvider, mailSender, cfg.EmailFromAddress)

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
