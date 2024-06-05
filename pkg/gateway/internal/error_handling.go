package internal

import (
	"log"
	"net/http"
)

func HandleInternalError(w http.ResponseWriter, e error) {
	w.WriteHeader(500)

	_, err := w.Write([]byte("Internal Error: " + e.Error()))
	if err != nil {
		log.Println("Error writing: " + err.Error())
	}
}
