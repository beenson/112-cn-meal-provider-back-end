package mail

import (
	"bytes"
	"strings"
	"time"
)

type Mail struct {
	From    string
	ReplyTo string
	To      []string
	Cc      []string
	Bcc     []string

	Subject string
	Body    string
}

func (m *Mail) Bytes() []byte {
	buf := bytes.Buffer{}

	buf.WriteString("From: " + m.From + "\r\n")
	buf.WriteString("Date: " + time.Now().Format(time.RFC1123Z) + "\r\n")

	if m.ReplyTo != "" {
		buf.WriteString("Reply-To: " + m.ReplyTo + "\r\n")
	}

	if len(m.To) > 0 {
		buf.WriteString("To: " + strings.Join(m.To, ", ") + "\r\n")
	}

	if len(m.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.Cc, ", ") + "\r\n")
	}

	if len(m.Bcc) > 0 {
		buf.WriteString("Bcc: " + strings.Join(m.Bcc, ", ") + "\r\n")
	}

	if m.Subject != "" {
		buf.WriteString("Subject: " + m.Subject + "\r\n")
	}

	buf.WriteString("\r\n")
	buf.WriteString(m.Body)

	return buf.Bytes()
}
