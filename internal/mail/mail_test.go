package mail

import (
	"bytes"
	"io"
	"net/mail"
	"strings"
	"testing"
)

func TestMail_Bytes(t *testing.T) {
	m := Mail{
		From:    "user1@example.com",
		ReplyTo: "reply1@example.com",

		To:  []string{"user2@example.com", "user3@example.com"},
		Cc:  []string{"user4@example.com", "user5@example.com"},
		Bcc: []string{"user6@example.com", "user7@example.com"},

		Subject: "Subject",
		Body:    "Body",
	}

	b := m.Bytes()

	message, err := mail.ReadMessage(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
		return
	}

	headers := message.Header
	if headers.Get("From") != m.From {
		t.Errorf("From header is wrong")
	}

	if headers.Get("Reply-To") != m.ReplyTo {
		t.Errorf("Reply-To header is wrong")
	}

	if headers.Get("To") != strings.Join(m.To, ", ") {
		t.Errorf("To header is wrong")
	}

	if headers.Get("Cc") != strings.Join(m.Cc, ", ") {
		t.Errorf("Cc header is wrong")
	}

	if headers.Get("Bcc") != strings.Join(m.Bcc, ", ") {
		t.Errorf("Bcc header is wrong")
	}

	if headers.Get("Subject") != m.Subject {
		t.Errorf("Subject header is wrong")
	}

	bodyBuf, err := io.ReadAll(message.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(bodyBuf) != m.Body {
		t.Errorf("Body header is wrong")
	}
}
