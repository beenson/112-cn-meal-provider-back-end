package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api"
	notify "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/notification"
)

func Mail(w http.ResponseWriter, r *http.Request) {
	if !auth(&w, r, "admin") {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	var msg api.PaymentNotifyMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		sendJSONError(w, "Could not parse json body !", http.StatusInternalServerError)
		return
	}
	conn, err := newClient(notificationTarget)
	if err != nil {
		sendJSONError(w, "Could not connect to notification service!", http.StatusInternalServerError)
		return
	}
	client := notify.NewNotificationServiceClient(conn)
	_, err = client.SendPayPaymentNotification(context.Background(), &notify.SendPayPaymentNotificationRequest{
		UserId: msg.Uid,
		Amount: msg.Amount,
	})
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK!"})
}
