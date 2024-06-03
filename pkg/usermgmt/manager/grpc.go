package manager

import (
	context "context"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"time"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/usermgmt/model"

	"github.com/jamesruan/sodium"
	common "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/common"
	protocol "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/user_mgmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (m *GRPCManager) PreAuth(ctx context.Context, user *protocol.UserID) (*protocol.Challenge, error) {
	defer fmt.Println("Get PreAuth Request")
	defer fmt.Println("pre-auth from " + *user.Id)

	m.db.Where("Uid = ?", user.Id).Delete(model.UserAuth{})

	// generate a random value as challenge
	seed := sodium.SignSeed{Bytes: []byte(*user.Id)}
	sodium.Randomize(&seed)

	// store challenge into db
	userAuth := model.UserAuth{
		Uid:       *user.Id,
		Challenge: b64.StdEncoding.EncodeToString(seed.Bytes),
	}
	if res := m.db.Create(&userAuth); res.Error != nil {
		defer fmt.Println("db.Create() failed: userAuth entry")
		return &protocol.Challenge{
			Value: seed.Bytes,
		}, status.Errorf(codes.Aborted, "Failed to push userAuth entry into database")
	}

	return &protocol.Challenge{
		Value: seed.Bytes,
	}, nil
}

func (m *GRPCManager) Auth(ctx context.Context, req *protocol.AuthRequest) (*common.AuthToken, error) {

	// get challenge from database
	var userAuth model.UserAuth
	m.db.Where("Uid = ?", req.Id.Id).Find(&userAuth)
	challenge, err := b64.StdEncoding.DecodeString(string(userAuth.Challenge))
	if err != nil {
		fmt.Println("Error decoding string:", err)
	}

	// get the user's public key
	fmt.Println("Get public key from database " + *req.Id.Id)
	var user []model.User
	m.db.Where("Uid = ?", req.Id.Id).Find(&user)
	if len(user) == 0 {
		fmt.Println("User record with name " + *req.Id.Id + " not found")
		return nil, status.Errorf(codes.NotFound, "Find no auth record")
	}
	var pk []byte
	if pk, err = b64.StdEncoding.DecodeString(user[0].PublicKey); err != nil {
		fmt.Println("Error decoding string:", err)
		return nil, status.Errorf(codes.Internal, "Error decoding string: %s", err)
	}

	// check the challenge with authRequest's signature
	err = sodium.Bytes(challenge).SignVerifyDetached(sodium.Signature{Bytes: sodium.Bytes(req.Response.Value)}, sodium.SignPublicKey{Bytes: sodium.Bytes(pk)})
	if err != nil {
		fmt.Println(*req.Id.Id + " fail the authentication with wrong signature")
		return nil, status.Errorf(codes.Aborted, "Challenge failed")
	} else {
		fmt.Println(*req.Id.Id + " success the authentication, return a AuthToken")
	}

	// generate a random token
	authToken := new([32]byte)
	rand.Read(authToken[:])

	userAuth.AuthAt = time.Now()
	userAuth.AuthToken = b64.StdEncoding.EncodeToString(authToken[:])
	m.db.Where("Uid = ?", req.Id.Id).Updates(userAuth)

	return &common.AuthToken{
		TokenValue: authToken[:],
	}, nil
}

func (m *GRPCManager) RegisterUser(ctx context.Context, req *protocol.RegisterUserRequest) (*common.Status, error) {
	// TODO: Auth admin user to do the operation
	user := model.User{
		Uid:        *req.Info.Info.Id,
		Group:      *req.Info.Info.Group,
		Username:   *req.Info.Info.Name,
		Department: *req.Info.Info.Department,
		Email:      *req.Info.Info.Email,
		PublicKey:  b64.StdEncoding.EncodeToString([]byte(*req.Info.PublicKey)),
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
	if res := m.db.Where("uid IN ?", req.UList.Uid).Find(&users); res.Error != nil {
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
