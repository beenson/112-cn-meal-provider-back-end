package notification

import (
	"context"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/sender/email"
	"testing"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/userinfo"
)

func TestNotificationService_SendPayPaymentNotification(t *testing.T) {
	sender := email.NewFakeEmailSender()
	userProvider := userinfo.NewFakeProvider()

	svc := NewService(userProvider, sender, "no-reply@example.com")

	err := svc.SendPayPaymentNotification(context.Background(), "1234", 30)
	if err != nil {
		return
	}

	if len(sender.SentMails) != 1 {
		t.Errorf("len(sender.SentMails) = %d, want %d", len(sender.SentMails), 1)
	}

	t.Logf("Mail body: %s", sender.SentMails[0].Body)
}
