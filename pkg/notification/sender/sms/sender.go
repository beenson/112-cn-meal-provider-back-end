package sms

type Sender interface {
	Send(to string, message string) error
}

type FakeSender struct {
}

func (f *FakeSender) Send(to string, message string) error {
	//TODO implement me
	panic("implement me")
}

func NewFakeSender() *FakeSender {
	return &FakeSender{}
}
