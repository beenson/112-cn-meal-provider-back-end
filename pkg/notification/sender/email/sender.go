package email

import "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal/mail"

type Sender interface {
	Send(mail mail.Mail) error
}

type FakeEmailSender struct {
	SentMails []mail.Mail
}

func (f *FakeEmailSender) Send(mail mail.Mail) error {
	f.SentMails = append(f.SentMails, mail)
	return nil
}

func NewFakeEmailSender() *FakeEmailSender {
	return &FakeEmailSender{}
}
