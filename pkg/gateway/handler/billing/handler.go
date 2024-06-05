package billing

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing"
	billingGRPC "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/transport/grpc"
)

type Handler struct {
	service billing.Service
}

func NewHandler(target string) *Handler {
	billingSvc, err := billingGRPC.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return &Handler{
		service: billingSvc,
	}
}
