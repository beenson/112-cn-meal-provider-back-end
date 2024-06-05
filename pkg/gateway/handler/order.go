package handler

import (
	"net/http"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func PostOrder(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func PutOrder(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
