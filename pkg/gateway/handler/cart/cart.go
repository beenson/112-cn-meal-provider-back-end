package cart

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	resp, err := h.client.GetCartItems(context.Background(), &order.GetCartItemsRequest{
		UserId: userId,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
func (h *Handler) PostCart(w http.ResponseWriter, r *http.Request) {

	var msg api.AddToCartRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.AddToCart(context.Background(), &order.AddToCartRequest{
		UserId: msg.UserId,
		FoodId: &order.FoodID{
			Id: msg.FoodId,
		},
		Amount: msg.Amount,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
func (h *Handler) PutCart(w http.ResponseWriter, r *http.Request) {
	var msg api.UpdateCartAmountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.UpdateCartAmount(context.Background(), &order.UpdateCartAmountRequest{
		Id: &order.CartItemID{
			Id: msg.Itemid,
		},
		Amount: msg.Amount,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
func (h *Handler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	var msg api.ClearCartRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.ClearCart(context.Background(), &order.ClearCartRequest{
		UserId: msg.UserId,
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
