package userinfo

import "context"

type fakeProvider struct{}

func NewFakeProvider() ProviderService {
	return &fakeProvider{}
}

func (f *fakeProvider) GetUser(ctx context.Context, userId string) (*User, error) {
	return &User{
		Name:  "John",
		Email: "john@example.com",
	}, nil
}
