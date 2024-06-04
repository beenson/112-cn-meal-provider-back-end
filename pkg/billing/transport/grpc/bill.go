package grpc

import (
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	pb "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/api/billing"
	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
)

func gRPCBillToModelBill(bill *pb.Bill) (*model.Bill, error) {
	id, err := strconv.Atoi(bill.Id.Id)
	if err != nil {
		return nil, err
	}

	return &model.Bill{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: bill.CreatedAt.AsTime(),
			UpdatedAt: bill.UpdatedAt.AsTime(),
		},
		UserId:  bill.UserId,
		OrderId: bill.OrderId,
		Amount:  int(bill.Amount),
	}, nil
}

func modelBillTogRPCBill(bill *model.Bill) (*pb.Bill, error) {
	return &pb.Bill{
		Id:        &pb.BillID{Id: strconv.Itoa(int(bill.ID))},
		CreatedAt: timestamppb.New(bill.CreatedAt),
		UpdatedAt: timestamppb.New(bill.UpdatedAt),

		UserId:  bill.UserId,
		OrderId: bill.OrderId,
		Amount:  int32(bill.Amount),
	}, nil
}
