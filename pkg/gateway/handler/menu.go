package handler

import (
	"net/http"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func PostMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func PutMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
}
