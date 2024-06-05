package order

import (
	"context"
	"encoding/json"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var msg api.GetOrderRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.GetOrder(context.Background(), &order.GetOrderRequest{
		Id: &order.OrderID{
			Id: msg.OrderId,
		},
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	resp, err := h.client.GetOrders(context.Background(), &order.GetOrdersRequest{
		UserId: userId,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
func (h *Handler) PostOrder(w http.ResponseWriter, r *http.Request) {
	var msg api.MakeOrderRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.MakeOrder(context.Background(), &order.MakeOrderRequest{
		UserId: msg.Uid,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
