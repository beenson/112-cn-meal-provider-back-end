package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	user "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/user_mgmt"
)

type Claims struct {
	Uid   string `json:"uid"`
	Group string `json:"group"`
	jwt.StandardClaims
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var msg api.UserLoginMsg

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	//conn, err := handler.newClient(handler.userTarget)
	//if err != nil {
	//	sendJSONError(w, "Could not connect to user management service!", http.StatusInternalServerError)
	//	return
	//}
	//client := user.NewUserManagementServiceClient(conn)
	resp, err := h.client.AuthUserLogin(r.Context(), &user.AuthUserRequest{
		Email:    &msg.Email,
		Password: &msg.Password,
	})
	if err != nil {
		internal.SendJSONError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Uid:   *resp.Uinfo[0].Id,
		Group: *resp.Uinfo[0].Group,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtKey)
	if err != nil {
		internal.SendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	internal.SendJSONResponse(w, struct {
		Token string `json:"token"`
	}{tokenString})
}

func Wrapper(group string, jwtKey string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			internal.SendJSONError(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			internal.SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if group != "" && claims.Group != group {
			internal.SendJSONError(w, "Not Valid User", http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
	}
}
