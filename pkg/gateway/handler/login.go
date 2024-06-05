package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamesruan/sodium"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	user "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
)

type Claims struct {
	Uid   string `json:"uid"`
	Group string `json:"group"`
	jwt.StandardClaims
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var NotInit bool = true
var jwtKey []byte

func Login(w http.ResponseWriter, r *http.Request) {
	var msg api.UserLoginMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(userTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to user management service!", http.StatusInternalServerError)
		return
	}
	client := user.NewUserManagementServiceClient(conn)
	resp, err := client.AuthUserLogin(context.Background(), &user.AuthUserRequest{
		Email:    &msg.Email,
		Password: &msg.Password,
	})
	if err != nil {
		sendJSONError(w, "Invalid email or password", http.StatusUnauthorized)
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
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func auth(w *http.ResponseWriter, r *http.Request, group string) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		sendJSONError(*w, "Missing token", http.StatusUnauthorized)
		return false
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		sendJSONError(*w, "Invalid token", http.StatusUnauthorized)
		return false
	}

	if group != "" && claims.Group != group {
		sendJSONError(*w, "Not Valid User", http.StatusUnauthorized)
		return false
	}

	return true
}

func sendJSONError(w http.ResponseWriter, errMessage string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := ErrorResponse{
		Error:   http.StatusText(status),
		Message: errMessage,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
