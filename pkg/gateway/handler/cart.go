package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.GetCartItemRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(orderTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
		return
	}
	client := order.NewOrderingServiceClient(conn)
	resp, err := client.GetCartItems(context.Background(), &order.GetCartItemsRequest{
		UserId: msg.UserId,
	})
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn.Close()
	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
func PostCart(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.AddToCartRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(orderTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
		return
	}
	client := order.NewOrderingServiceClient(conn)
	resp, err := client.AddToCart(context.Background(), &order.AddToCartRequest{
		UserId: msg.UserId,
		FoodId: &order.FoodID{
			Id: msg.FoodId,
		},
		Amount: msg.Amount,
	})
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn.Close()
	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
func PutCart(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.UpdateCartAmountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(orderTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
		return
	}
	client := order.NewOrderingServiceClient(conn)
	resp, err := client.UpdateCartAmount(context.Background(), &order.UpdateCartAmountRequest{
		Id: &order.CartItemID{
			Id: msg.Itemid,
		},
		Amount: msg.Amount,
	})
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn.Close()
	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.ClearCartRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(orderTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to order service!", http.StatusInternalServerError)
		return
	}
	client := order.NewOrderingServiceClient(conn)
	resp, err := client.ClearCart(context.Background(), &order.ClearCartRequest{
		UserId: msg.UserId,
	})
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn.Close()
	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
