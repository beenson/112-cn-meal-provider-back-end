package billing

import (
	"net/http"
	"strconv"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
)

type HTTPBill struct {
	Id      string `json:"id"`
	UserId  string `json:"userId"`
	OrderId string `json:"orderId"`
	Amount  int    `json:"amount"`
}

func modelBillToHTTPBill(bill *model.Bill) *HTTPBill {
	return &HTTPBill{
		Id:      strconv.Itoa(int(bill.ID)),
		UserId:  bill.UserId,
		OrderId: bill.OrderId,
		Amount:  bill.Amount,
	}
}

func (h *Handler) GetBills(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := r.URL.Query().Get("user_id")
	bills, err := h.service.GetBills(ctx, userId)
	if err != nil {
		internal.HandleInternalError(w, err)
		return
	}

	data := make([]*HTTPBill, 0)
	for _, bill := range bills {
		data = append(data, modelBillToHTTPBill(&bill))
	}

	resp := struct {
		Bills []*HTTPBill `json:"bills"`
	}{data}

	w.WriteHeader(http.StatusOK)
	internal.SendJSONResponse(w, resp)
}
