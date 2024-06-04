package main

import (
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/sender/email"
	"net/smtp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification"
	grpcTransport "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/transport/grpc"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/userinfo"
)

func main() {
	userProvider, err := userinfo.NewProvider(
		"localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	mailSender := email.NewSMTPSender("",
		smtp.PlainAuth("", "", "", ""),
	)

	svc := notification.NewService(userProvider, mailSender, "")

	maxSize := 256 << 20
	grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

	server, err := grpcTransport.NewServer("0.0.0.0:50052", grpcOpts, svc)
	if err != nil {
		return
	}

	server.Serve()

}
