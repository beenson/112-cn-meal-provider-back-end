package main

import (
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler/order"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/gorilla/mux"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler/auth"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler/cart"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/handler/food"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	// ping
	r.HandleFunc("/ping", handler.Ping)

	// auth
	{
		authHandler := auth.NewHandler(cfg.Service.UserMgmtTarget, cfg.JwtSecret)
		// login, {"email": string, "password": string}
		r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	}

	// billing
	{
		billingHandler := billing.NewHandler(cfg.Service.BillingTarget)
		// get bills, query param `user_id`
		r.HandleFunc("/api/billing/bill", billingHandler.GetBills).Methods(http.MethodGet)
		// get payments, query param `user_id`
		r.HandleFunc("/api/billing/payment", billingHandler.GetPayments).Methods(http.MethodGet)
		// create payments, {"userId": string, "amount": number}
		r.HandleFunc("/api/billing/payment", billingHandler.PostPayments).Methods(http.MethodPost)
	}

	// menu
	foodHandler := food.NewHandler(cfg.Service.OrderingTarget)
	//r.HandleFunc("/api/food", auth.Wrapper("", cfg.JwtSecret, foodHandler.GetMenu)).Methods(http.MethodGet) // user + admin
	//r.HandleFunc("/api/food", auth.Wrapper("admin", cfg.JwtSecret, foodHandler.PostMenu)).Methods(http.MethodPost)
	//r.HandleFunc("/api/food", auth.Wrapper("admin", cfg.JwtSecret, foodHandler.PutMenu)).Methods(http.MethodPut)
	//r.HandleFunc("/api/food", auth.Wrapper("admin", cfg.JwtSecret, foodHandler.DeleteMenu)).Methods(http.MethodDelete)
	r.HandleFunc("/api/food", foodHandler.GetMenu).Methods(http.MethodGet) // user + admin
	r.HandleFunc("/api/food", foodHandler.PostMenu).Methods(http.MethodPost)
	r.HandleFunc("/api/food", foodHandler.PutMenu).Methods(http.MethodPut)
	r.HandleFunc("/api/food", foodHandler.DeleteMenu).Methods(http.MethodDelete)

	// cart
	cartHandler := cart.NewHandler(cfg.Service.OrderingTarget)
	r.HandleFunc("/api/cart", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/api/cart", cartHandler.PostCart).Methods("POST")

	// order
	orderHandler := order.NewHandler(cfg.Service.OrderingTarget)
	r.HandleFunc("/api/order", orderHandler.GetOrders).Methods("GET")
	r.HandleFunc("/api/order", orderHandler.PostOrder).Methods("POST")

	// order user
	//r.HandleFunc("/api/order", handler.GetOrder).Methods("GET") //user + admin
	//r.HandleFunc("/api/orders", handler.GetOrders).Methods("GET")
	//r.HandleFunc("/api/order", handler.PostOrder).Methods("POST")
	////mail
	//r.HandleFunc("/api/mail", handler.Mail)
	//
	////comment user
	//r.HandleFunc("/api/comment", handler.GetComment).Methods("GET")
	//r.HandleFunc("/api/comment", handler.PostComment).Methods("POST")
	//r.HandleFunc("/api/comment", handler.PutComment).Methods("PUT")
	//r.HandleFunc("/api/comment", handler.DeleteComment).Methods("DELELTE")

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)
	r.Use(logMiddleware)

	log.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(cfg.Address, r))
}
