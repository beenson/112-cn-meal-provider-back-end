package notification

import (
	"context"
	"errors"
	"fmt"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal/mail"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/sender/email"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/notification/userinfo"
)

type Service interface {
	SendPayPaymentNotification(ctx context.Context, userId string, amountToPay int) error
}

type notificationService struct {
	userProvider userinfo.ProviderService
	emailSender  email.Sender

	fromAddress string
}

func (n *notificationService) SendPayPaymentNotification(ctx context.Context, userId string, amountToPay int) error {
	user, err := n.userProvider.GetUser(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	builder := mail.NewBuilder(n.fromAddress)
	builder.AddRecipients(user.Email)
	builder.SetSubject("Payment Request from Food Provider")
	builder.SetBody(
		fmt.Sprintf("Dear %s,\n\nYou need to pay $%d this month, please make the payment ASAP.", user.Name, amountToPay),
	)

	err = n.emailSender.Send(builder.Build())
	if err != nil {
		return err
	}

	return nil
}

func NewService(userProvider userinfo.ProviderService, emailSender email.Sender, fromAddress string) Service {
	return &notificationService{
		userProvider: userProvider,
		emailSender:  emailSender,

		fromAddress: fromAddress,
	}
}
