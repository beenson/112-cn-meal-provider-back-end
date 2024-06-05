package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	rating "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/rating"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
	// if !auth(&w, r, "") {
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// 	var msg api.ReadCommentMsg
	// 	decoder := json.NewDecoder(r.Body)
	// 	err := decoder.Decode(&msg)
	// 	if err != nil {
	// 		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	fmt.Printf("Received payload: %+v\n", msg)
	// 	conn, err := newClient(ratingTarget)
	// 	if err != nil {
	// 		sendJSONError(w, "Could not connect to rating service!", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	client := rating.NewFoodRatingServiceClient(conn)
	// 	if msg.CommentId != 0 {
	// 		// 	resp, err := client.ReadFeedbackByCommentId(context.Background(), &rating.CommentId{
	// 		// 		Id: &msg.CommentId,
	// 		// 	})
	// 		// } else if msg.FoodId != 0 {
	// 		// 	if msg.Rate != 0 {
	// 		// 		resp, err := client.ReadFeedbackByFoodRating(context.Background(), &rating.CommentId{
	// 		// 			Id: &msg.CommentId,
	// 		// 		})
	// 		// 	}

	// 	}
	// 	if err != nil {
	// 		sendJSONError(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	conn.Close()
	// 	jsonBytes, err := protojson.Marshal(resp)
	// 	if err != nil {
	// 		sendJSONError(w, err.Error(), http.StatusBadRequest)
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(jsonBytes)
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.CreateCommentMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(ratingTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to rating service!", http.StatusInternalServerError)
		return
	}
	client := rating.NewFoodRatingServiceClient(conn)
	resp, err := client.CreateFeedback(context.Background(), &rating.Feedback{
		FoodId:  &msg.FoodId,
		Uid:     &msg.Uid,
		Rating:  &msg.Rate,
		Comment: &msg.Comment,
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
func PutComment(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.UpdateCommentMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(ratingTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to rating service!", http.StatusInternalServerError)
		return
	}
	client := rating.NewFoodRatingServiceClient(conn)
	resp, err := client.UpdateFeedback(context.Background(), &rating.UpdateRequest{
		Id: &rating.CommentId{
			Id: &msg.CommentId,
		},
		Fb: &rating.Feedback{
			FoodId:  &msg.FoodId,
			Uid:     &msg.Uid,
			Rating:  &msg.Rate,
			Comment: &msg.Comment,
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
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.DeleteCommentMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(ratingTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to rating service!", http.StatusInternalServerError)
		return
	}
	client := rating.NewFoodRatingServiceClient(conn)
	resp, err := client.DeleteFeedback(context.Background(), &rating.CommentId{
		Id: &msg.CommentId,
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
