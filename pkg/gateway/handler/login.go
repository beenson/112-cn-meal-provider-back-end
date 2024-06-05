package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamesruan/sodium"

	user "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
)

type UserLoginMsg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Uid   string `json:"uid"`
	Group string `json:"group"`
	jwt.StandardClaims
}

var NotInit bool = true
var jwtKey []byte

func Login(w http.ResponseWriter, r *http.Request) {
	var msg UserLoginMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		http.Error(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(userTarget)
	if err != nil {
		http.Error(w, "Could not connect to user management service!", http.StatusInternalServerError)
		return
	}
	client := user.NewUserManagementServiceClient(conn)
	resp, err := client.AuthUserLogin(context.Background(), &user.AuthUserRequest{
		Email:    &msg.Email,
		Password: &msg.Password,
	})
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	conn.Close()

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Uid:   *resp.Uinfo[0].Id,
		Group: *resp.Uinfo[0].Group,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	if NotInit {
		NotInit = false
		seed := sodium.SignSeed{}
		sodium.Randomize(&seed)
		jwtKey = seed.Bytes
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
