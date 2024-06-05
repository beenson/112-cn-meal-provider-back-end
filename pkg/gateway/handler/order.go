package handler

//import (
//	"context"
//	"encoding/json"
//	"net/http"
//
//	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
//	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
//	"google.golang.org/protobuf/encoding/protojson"
//)
//
//func GetOrder(w http.ResponseWriter, r *http.Request) {
//	if !auth(&w, r, "") {
//		return
//	}
//	// w.Header().Set("Content-Type", "application/json")
//	var msg api.GetOrderRequest
//	decoder := json.NewDecoder(r.Body)
//	err := decoder.Decode(&msg)
//	if err != nil {
//		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
//		return
//	}
//	conn, err := newClient(orderTarget)
//	if err != nil {
//		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
//		return
//	}
//	client := order.NewOrderingServiceClient(conn)
//	resp, err := client.GetOrder(context.Background(), &order.GetOrderRequest{
//		Id: &order.OrderID{
//			Id: msg.OrderId,
//		},
//	})
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	conn.Close()
//	jsonBytes, err := protojson.Marshal(resp)
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonBytes)
//}
//func GetOrders(w http.ResponseWriter, r *http.Request) {
//	if !auth(&w, r, "") {
//		return
//	}
//	// w.Header().Set("Content-Type", "application/json")
//	var msg api.GetOrdersRequest
//	decoder := json.NewDecoder(r.Body)
//	err := decoder.Decode(&msg)
//	if err != nil {
//		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
//		return
//	}
//	conn, err := newClient(orderTarget)
//	if err != nil {
//		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
//		return
//	}
//	client := order.NewOrderingServiceClient(conn)
//	resp, err := client.GetOrders(context.Background(), &order.GetOrdersRequest{
//		UserId: msg.Uid,
//	})
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	conn.Close()
//	jsonBytes, err := protojson.Marshal(resp)
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonBytes)
//}
//func PostOrder(w http.ResponseWriter, r *http.Request) {
//	if !auth(&w, r, "") {
//		return
//	}
//	// w.Header().Set("Content-Type", "application/json")
//	var msg api.MakeOrderRequest
//	decoder := json.NewDecoder(r.Body)
//	err := decoder.Decode(&msg)
//	if err != nil {
//		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
//		return
//	}
//	conn, err := newClient(orderTarget)
//	if err != nil {
//		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
//		return
//	}
//	client := order.NewOrderingServiceClient(conn)
//	resp, err := client.MakeOrder(context.Background(), &order.MakeOrderRequest{
//		UserId: msg.Uid,
//	})
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	conn.Close()
//	jsonBytes, err := protojson.Marshal(resp)
//	if err != nil {
//		sendJSONError(w, err.Error(), http.StatusBadRequest)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonBytes)
//}
