package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	order "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.GetFoodRequest
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
	if msg.Id == "" {
		resp, err := client.GetFoods(context.Background(), &order.GetFoodsRequest{})
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
	} else {
		resp, err := client.GetFood(context.Background(), &order.GetFoodRequest{
			Id: &order.FoodID{
				Id: msg.Id,
			},
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
}

func PostMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.CreateFoodRequest
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
	resp, err := client.CreateFood(context.Background(), &order.CreateFoodRequest{
		Input: &order.CreateFoodInput{
			Name:        msg.Name,
			Description: msg.Description,
			Price:       msg.Price,
			ImageUrl:    msg.URL,
		},
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

func PutMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.UpdateFoodRequest
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
	resp, err := client.UpdateFood(context.Background(), &order.UpdateFoodRequest{
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
func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.DeleteFoodRequest
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
	resp, err := client.DeleteFood(context.Background(), &order.DeleteFoodRequest{
		Id: &order.FoodID{
			Id: msg.Id,
		},
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
