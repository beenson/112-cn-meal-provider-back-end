package handler

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	// authHeader := r.Header.Get("Authorization")
	// if authHeader == "" {
	// 	http.Error(w, "Missing token", http.StatusUnauthorized)
	// 	return
	// }

	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// claims := &Claims{}

	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })

	// if err != nil || !token.Valid {
	// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
	// 	return
	// }
	// username := r.Context().Value("Uid").(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"msg": "pong!"})
}
