package test

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	db2 "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal/db"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/sender/email"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/userinfo"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/manager"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/model"
)

func Test_NotificationAndUserMgmt(t *testing.T) {
	{
		db, err := db2.Config{
			Type: "sqlite",
			DSN:  "file::memory:",
		}.GetGormDB()
		if err != nil {
			t.Fatal(err)
		}

		err = db.AutoMigrate(model.GetModels()...)
		if err != nil {
			t.Fatal(err)
		}

		db.Create(&model.User{
			Uid:        "1234",
			Group:      "user",
			Username:   "user1234",
			Department: "dept A",
			Email:      "user1234@gmail.com",
			Password:   "",
		})

		maxSize := 256 << 20
		grpcOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize)}

		m := manager.NewManager(":8000", grpcOpts, db)
		go m.Serve()
	}

	userProvider, err := userinfo.NewProvider(
		"localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	mailSender := email.NewFakeEmailSender()
	svc := notification.NewService(userProvider, mailSender, "no-reply@example.com")

	t.Run("userProvider should get user info", func(t *testing.T) {
		user, err := userProvider.GetUser(context.Background(), "1234")
		if err != nil {
			t.Fatal(err)
		}

		if user == nil {
			t.Fatal("user is nil")
		}

		if user.Name != "user1234" {
			t.Fatal("user name is wrong")
		}

		if user.Email != "user1234@gmail.com" {
			t.Fatal("user email is wrong")
		}
	})

	t.Run("", func(t *testing.T) {
		err := svc.SendPayPaymentNotification(context.Background(), "1234", 30)
		if err != nil {
			t.Fatal(err)
		}

		if len(mailSender.SentMails) != 1 {
			t.Fatal("len(mailSender.SentMails) != 1")
		}

		t.Logf("%v", mailSender.SentMails[0])
	})
}
