package handler

import (
	"net/http"
)

func GetBill(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func PostBill(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
