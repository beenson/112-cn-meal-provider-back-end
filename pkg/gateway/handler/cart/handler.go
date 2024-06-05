package cart

import (
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/ordering"
	"log"
)

type Handler struct {
	client ordering.OrderingServiceClient
}

func NewHandler(target string) *Handler {
	conn, err := internal.NewClient(target)
	if err != nil {
		log.Fatal(err)
	}

	client := ordering.NewOrderingServiceClient(conn)

	return &Handler{
		client: client,
	}
}
