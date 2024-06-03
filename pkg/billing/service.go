package billing

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
)

type Service interface {
	CreateBill(userId string, orderId string, amount int) (*model.Bill, error)
	CreatePayment(userId string, amount int) (*model.Payment, error)

	GetBill(id uint) (*model.Bill, error)
	GetBills(userId string) ([]model.Bill, error)

	GetPayment(id uint) (*model.Payment, error)
	GetPayments(userId string) ([]model.Payment, error)
}

type billingService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &billingService{db}
}

func (svc *billingService) CreateBill(userId string, orderId string, amount int) (*model.Bill, error) {
	b := model.NewBill(userId, orderId, amount)

	result := svc.db.Clauses(clause.Returning{}).Create(&b)
	if result.Error != nil {
		return nil, result.Error
	}

	return &b, nil
}

func (svc *billingService) CreatePayment(userId string, amount int) (*model.Payment, error) {
	p := model.NewPayment(userId, amount)

	result := svc.db.Clauses(clause.Returning{}).Create(&p)
	if result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

func (svc *billingService) GetBill(id uint) (*model.Bill, error) {
	var b model.Bill

	svc.db.First(&b, id)
	if b.Model.ID == 0 {
		return nil, nil
	}

	return &b, nil
}

func (svc *billingService) GetBills(userId string) ([]model.Bill, error) {
	var bills []model.Bill

	svc.db.Where("user_id = ?", userId).Find(&bills)

	return bills, nil
}

func (svc *billingService) GetPayment(id uint) (*model.Payment, error) {
	var p model.Payment

	svc.db.First(&p, id)
	if p.Model.ID == 0 {
		return nil, nil
	}

	return &p, nil
}

func (svc *billingService) GetPayments(userId string) ([]model.Payment, error) {
	var payments []model.Payment

	svc.db.Where("user_id = ?", userId).Find(&payments)

	return payments, nil
}
