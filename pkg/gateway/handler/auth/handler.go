package auth

import (
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
	"log"
)

type Handler struct {
	client user_mgmt.UserManagementServiceClient
	jwtKey []byte
}

func NewHandler(target string, jwtKey string) *Handler {
	conn, err := internal.NewClient(target)
	if err != nil {
		log.Fatal(err)
	}

	client := user_mgmt.NewUserManagementServiceClient(conn)

	return &Handler{
		client: client,
		jwtKey: []byte(jwtKey),
	}
}
