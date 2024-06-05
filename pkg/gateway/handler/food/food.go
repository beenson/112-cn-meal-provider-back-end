package food

import (
	"context"
	"encoding/json"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
)

func (h *Handler) GetMenu(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		resp, err := h.client.GetFoods(context.Background(), &order.GetFoodsRequest{})
		if err != nil {
			internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		internal.SendProtoJSONResponse(w, resp)
	} else {
		resp, err := h.client.GetFood(context.Background(), &order.GetFoodRequest{
			Id: &order.FoodID{
				Id: userId,
			},
		})
		if err != nil {
			internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		internal.SendProtoJSONResponse(w, resp)
	}
}

func (h *Handler) PostMenu(w http.ResponseWriter, r *http.Request) {
	var msg api.CreateFoodRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.CreateFood(context.Background(), &order.CreateFoodRequest{
		Input: &order.CreateFoodInput{
			Name:        msg.Name,
			Description: msg.Description,
			Price:       msg.Price,
			ImageUrl:    msg.URL,
		},
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}

func (h *Handler) PutMenu(w http.ResponseWriter, r *http.Request) {
	var msg api.UpdateFoodRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.UpdateFood(context.Background(), &order.UpdateFoodRequest{
		Id: &order.FoodID{
			Id: msg.Id,
		},
		Input: &order.UpdateFoodInput{
			Name:        &msg.Name,
			Description: &msg.Description,
			Price:       &msg.Price,
			ImageUrl:    &msg.URL,
		},
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
func (h *Handler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	var msg api.DeleteFoodRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		internal.SendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}

	resp, err := h.client.DeleteFood(context.Background(), &order.DeleteFoodRequest{
		Id: &order.FoodID{
			Id: msg.Id,
		},
	})
	if err != nil {
		internal.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.SendProtoJSONResponse(w, resp)
}
