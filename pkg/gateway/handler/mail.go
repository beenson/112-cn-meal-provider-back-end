package handler

import (
	"net/http"
)

func Mail(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
	// w.Header().Set("Content-Type", "application/json")
}