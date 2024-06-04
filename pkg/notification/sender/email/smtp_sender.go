package email

import (
	"net/smtp"
	"slices"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal/mail"
)

type SMTPSender struct {
	addr string
	auth smtp.Auth
}

func (s *SMTPSender) Send(mail mail.Mail) error {
	to := slices.Concat(mail.To, mail.Cc, mail.Bcc)
	msg := mail.Bytes()

	err := smtp.SendMail(s.addr, s.auth, mail.From, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func NewSMTPSender(addr string, auth smtp.Auth) *SMTPSender {
	return &SMTPSender{
		addr: addr,
		auth: auth,
	}
}
