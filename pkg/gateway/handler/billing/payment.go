package billing

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
)

type HTTPPayment struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Amount int    `json:"amount"`
}

func modelPaymentToHTTPPayment(payment *model.Payment) *HTTPPayment {
	return &HTTPPayment{
		Id:     strconv.Itoa(int(payment.ID)),
		UserId: payment.UserId,
		Amount: payment.Amount,
	}
}

func (h *Handler) GetPayments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := r.URL.Query().Get("user_id")
	payments, err := h.service.GetPayments(ctx, userId)
	if err != nil {
		internal.HandleInternalError(w, err)
		return
	}

	data := make([]*HTTPPayment, 0)
	for _, payment := range payments {
		data = append(data, modelPaymentToHTTPPayment(&payment))
	}

	resp := struct {
		Payments []*HTTPPayment `json:"payments"`
	}{data}

	w.WriteHeader(http.StatusOK)
	internal.SendJSONResponse(w, resp)
}

func (h *Handler) PostPayments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body := struct {
		UserId string `json:"userId"`
		Amount int    `json:"amount"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		internal.HandleInternalError(w, err)
		return
	}

	payment, err := h.service.CreatePayment(ctx, body.UserId, body.Amount)
	if err != nil {
		internal.HandleInternalError(w, err)
		return
	}

	p := modelPaymentToHTTPPayment(payment)

	resp := struct {
		Payment *HTTPPayment `json:"payment"`
	}{p}

	w.WriteHeader(http.StatusCreated)
	internal.SendJSONResponse(w, resp)
}
