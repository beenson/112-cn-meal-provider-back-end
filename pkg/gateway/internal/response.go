package internal

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

func SendProtoJSONResponse(w http.ResponseWriter, resp proto.Message) {
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		SendJSONError(w, err.Error(), http.StatusBadRequest)
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(err)
	}
}

func SendJSONResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Panicln(err)
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func SendJSONError(w http.ResponseWriter, errMessage string, status int) {
	errorResponse := ErrorResponse{
		Error:   http.StatusText(status),
		Message: errMessage,
	}

	w.WriteHeader(status)
	SendJSONResponse(w, errorResponse)
}
