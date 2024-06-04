package mail

type Builder struct {
	mail Mail
}

func NewBuilder(from string) *Builder {
	return &Builder{mail: Mail{
		From: from,
	}}
}

func (b *Builder) SetFrom(from string) *Builder {
	b.mail.From = from
	return b
}

func (b *Builder) SetReplyTo(replyTo string) *Builder {
	b.mail.ReplyTo = replyTo
	return b
}

func (b *Builder) AddRecipients(recipients ...string) *Builder {
	b.mail.To = append(b.mail.To, recipients...)
	return b
}

func (b *Builder) AddCcRecipients(cc ...string) *Builder {
	b.mail.Cc = append(b.mail.Cc, cc...)
	return b
}

func (b *Builder) AddBccRecipients(bcc ...string) *Builder {
	b.mail.Bcc = append(b.mail.Bcc, bcc...)
	return b
}

func (b *Builder) SetSubject(subject string) *Builder {
	b.mail.Subject = subject
	return b
}

func (b *Builder) SetBody(body string) *Builder {
	b.mail.Body = body
	return b
}

func (b *Builder) Build() Mail {
	return b.mail
}
