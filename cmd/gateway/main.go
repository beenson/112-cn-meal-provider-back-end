package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler"
)

func main() {
	r := mux.NewRouter()
	//menu admin
	r.HandleFunc("/api/menu", handler.GetMenu).Methods("GET") //user + admin
	r.HandleFunc("/api/menu", handler.PostMenu).Methods("POST")
	r.HandleFunc("/api/menu", handler.PutMenu).Methods("PUT")
	r.HandleFunc("/api/menu", handler.DeleteMenu).Methods("DELELTE")
	//order user
	r.HandleFunc("/api/order", handler.GetOrder).Methods("GET") //user + admin
	r.HandleFunc("/api/order", handler.PostOrder).Methods("POST")
	r.HandleFunc("/api/order", handler.PutOrder).Methods("PUT")
	r.HandleFunc("/api/order", handler.DeleteOrder).Methods("DELELTE")
	//mail
	r.HandleFunc("/api/mail", handler.Mail)
	//login
	r.HandleFunc("/login", handler.Login)
	//ping
	r.HandleFunc("/ping", handler.Ping)
	//comment user
	r.HandleFunc("/api/comment", handler.GetComment).Methods("GET")
	r.HandleFunc("/api/comment", handler.PostComment).Methods("POST")
	r.HandleFunc("/api/comment", handler.PutComment).Methods("PUT")
	r.HandleFunc("/api/comment", handler.DeleteComment).Methods("DELELTE")
	log.Fatal(http.ListenAndServe(":55688", r))
}
