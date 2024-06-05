package manager

import (
	context "context"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/model"

	common "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/common"
	protocol "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (m *GRPCManager) RegisterUser(ctx context.Context, req *protocol.RegisterUserRequest) (*common.Status, error) {
	// TODO: Auth admin user to do the operation
	user := model.User{
		Uid:        *req.Info.Info.Id,
		Group:      *req.Info.Info.Group,
		Username:   *req.Info.Info.Name,
		Department: *req.Info.Info.Department,
		Email:      *req.Info.Info.Email,
		Password:   *req.Info.Password,
	}
	m.db.Create(&user)
	code := int32(0)
	return &common.Status{
		Code: &code,
	}, nil
}

func (m *GRPCManager) UnregisterUser(ctx context.Context, req *protocol.UnregisterUserRequest) (*common.Status, error) {
	// TODO: Auth admin user to do the operation
	m.db.Where("Uid = ?", req.Target.Id).Delete(model.User{})
	m.db.Where("Uid = ?", req.Target.Id).Delete(model.UserAuth{})
	code := int32(0)
	return &common.Status{
		Code: &code,
	}, nil
}

func (m *GRPCManager) GetSysUsers(context.Context, *protocol.GetUserRequest) (*protocol.GetUserReply, error) {
	var reply protocol.GetUserReply
	var users []model.User
	if res := m.db.Find(&users); res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(users); i++ {
		reply.Uinfo = append(reply.Uinfo, &common.UserInfo{
			Id:         &users[i].Uid,
			Name:       &users[i].Username,
			Group:      &users[i].Group,
			Department: &users[i].Department,
			Email:      &users[i].Email,
		})
	}
	return &reply, nil
}

func (m *GRPCManager) GetUserByIds(ctx context.Context, req *protocol.GetUserRequest) (*protocol.GetUserReply, error) {
	var reply protocol.GetUserReply
	var users []model.User
	if res := m.db.Where("Uid IN ?", req.UList.Uid).Find(&users); res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(users); i++ {
		reply.Uinfo = append(reply.Uinfo, &common.UserInfo{
			Id:         &users[i].Uid,
			Name:       &users[i].Username,
			Group:      &users[i].Group,
			Department: &users[i].Department,
			Email:      &users[i].Email,
		})
	}
	return &reply, nil
}

func (m *GRPCManager) AuthUserLogin(ctx context.Context, req *protocol.AuthUserRequest) (*protocol.GetUserReply, error) {
	var reply protocol.GetUserReply
	var users []model.User
	if res := m.db.Where("Email = ?", req.Email).Find(&users); res.Error != nil {
		return nil, res.Error
	} else if len(users) == 0 {
		return nil, status.Error(codes.NotFound, "User Not Found")
	}

	if users[0].Password != *req.Password {
		return nil, status.Error(codes.Unauthenticated, "Password Not Matched!")
	}
	reply.Uinfo = append(reply.Uinfo, &common.UserInfo{
		Id:         &users[0].Uid,
		Name:       &users[0].Username,
		Group:      &users[0].Group,
		Department: &users[0].Department,
		Email:      &users[0].Email,
	})
	return &reply, nil
}
