package userinfo

import (
	"context"

	"google.golang.org/grpc"

	pbusermgmt "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
)

type User struct {
	Name  string
	Email string
}

type ProviderService interface {
	GetUser(ctx context.Context, userId string) (*User, error)
}

type provider struct {
	userMgmtClient pbusermgmt.UserManagementServiceClient
}

func (e *provider) GetUser(ctx context.Context, userId string) (*User, error) {
	request := &pbusermgmt.GetUserRequest{
		UList: &pbusermgmt.UserList{
			Uid: []string{userId},
		},
	}

	ids, err := e.userMgmtClient.GetUserByIds(ctx, request)
	if err != nil {
		return nil, err
	}

	if len(ids.Uinfo) == 0 {
		return nil, nil
	}

	return &User{
		Name:  ids.Uinfo[0].GetName(),
		Email: ids.Uinfo[0].GetEmail(),
	}, nil
}

func NewProvider(target string, opts ...grpc.DialOption) (ProviderService, error) {
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	return &provider{
		userMgmtClient: pbusermgmt.NewUserManagementServiceClient(conn),
	}, nil
}
