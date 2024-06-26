package grpc

import (
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/billing"
)

func gRPCPaymentToModelPayment(p *pb.Payment) (*model.Payment, error) {
	if p == nil {
		return nil, nil
	}

	id, err := strconv.Atoi(p.Id.Id)
	if err != nil {
		return nil, err
	}

	return &model.Payment{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: p.CreatedAt.AsTime(),
			UpdatedAt: p.UpdatedAt.AsTime(),
		},

		UserId: p.UserId,
		Amount: int(p.Amount),
	}, nil
}

func modelPaymentTogRPCPayment(p *model.Payment) (*pb.Payment, error) {
	if p == nil {
		return nil, nil
	}

	return &pb.Payment{
		Id:        &pb.PaymentID{Id: strconv.Itoa(int(p.ID))},
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),

		UserId: p.UserId,
		Amount: int32(p.Amount),
	}, nil
}
