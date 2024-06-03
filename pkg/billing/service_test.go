package billing

import (
	"errors"
	"os"
	"path"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/billing/model"
)

func setupService(t *testing.T) Service {
	dbFile := path.Join(t.TempDir(), "test.db")

	stat, err := os.Stat(dbFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	} else {
		if stat.IsDir() {
			panic("db file is dir")
		}

		err = os.Remove(dbFile)
		if err != nil {
			panic(err)
		}
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Bill{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Payment{})
	if err != nil {
		panic(err)
	}

	return NewService(db)
}

func cleanupService() {
}

func checkModelIsNotZero(t *testing.T, m *gorm.Model) {
	if m.CreatedAt.IsZero() {
		t.Fatal("m.CreatedAt is zero")
	}

	if m.UpdatedAt.IsZero() {
		t.Fatal("m.UpdatedAt is zero")
	}
}

func checkModelIsEqual(t *testing.T, m1, m2 *gorm.Model) {
	if m1.ID != m2.ID {
		t.Fatal("m1.ID != m2.ID")
	}

	if m1.CreatedAt != m2.CreatedAt {
		t.Fatal("m1.CreatedAt != m2.CreatedAt")
	}

	if m1.UpdatedAt != m2.UpdatedAt {
		t.Fatal("m1.UpdatedAt != m2.UpdatedAt")
	}

	if m1.DeletedAt != m2.DeletedAt {
		t.Fatal("m1.DeletedAt != m2.DeletedAt")
	}
}

func checkBillIsEqual(t *testing.T, b1, b2 *model.Bill) {
	checkModelIsEqual(t, &b1.Model, &b2.Model)

	if b1.UserId != b2.UserId {
		t.Fatal("b1.UserId != b2.UserId")
	}

	if b1.OrderId != b2.OrderId {
		t.Fatal("b1.OrderId != b2.OrderId")
	}

	if b1.Amount != b2.Amount {
		t.Fatal("b1.Amount != b2.Amount")
	}
}

func checkPaymentIsEqual(t *testing.T, p1, p2 *model.Payment) {
	checkModelIsEqual(t, &p1.Model, &p2.Model)

	if p1.UserId != p2.UserId {
		t.Fatal("p1.UserId != p2.UserId")
	}

	if p1.Amount != p2.Amount {
		t.Fatal("p1.Amount != p2.Amount")
	}
}

func TestBillingService_CreateBill(t *testing.T) {
	svc := setupService(t)
	defer cleanupService()

	userId := "1234"
	orderId := "abc"
	amount := 10

	bill, err := svc.CreateBill(userId, orderId, amount)
	if err != nil {
		t.Fatal(err)
	}

	checkModelIsNotZero(t, &bill.Model)

	if bill.UserId != userId {
		t.Fatal("bill.UserId is wrong")
	}

	if bill.OrderId != orderId {
		t.Fatal("bill.OrderId is wrong")
	}

	if bill.Amount != amount {
		t.Fatal("bill.Amount is wrong")
	}
}

func TestBillingService_CreatePayment(t *testing.T) {
	svc := setupService(t)
	defer cleanupService()

	userId := "1234"
	amount := 10

	payment, err := svc.CreatePayment(userId, amount)
	if err != nil {
		t.Fatal(err)
	}

	checkModelIsNotZero(t, &payment.Model)

	if payment.UserId != userId {
		t.Fatal("payment.UserId is wrong")
	}

	if payment.Amount != amount {
		t.Fatal("payment.Amount is wrong")
	}
}

func TestBillingService_GetBill(t *testing.T) {
	t.Run("should return bill", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		userId := "1234"
		orderId := "abc"
		amount := 10

		b1, err := svc.CreateBill(userId, orderId, amount)
		if err != nil {
			t.Fatal(err)
		}

		bill, err := svc.GetBill(b1.ID)
		if err != nil {
			t.Fatal(err)
		}

		checkModelIsNotZero(t, &bill.Model)

		checkBillIsEqual(t, bill, b1)
	})

	t.Run("should return nil when not found", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		b, err := svc.GetBill(1)
		if err != nil {
			t.Fatal(err)
		}

		if b != nil {
			t.Fatal("bill is not nil")
		}
	})
}

func TestBillingService_GetBills(t *testing.T) {
	t.Run("should return bills", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		userId := "1234"

		b1, err := svc.CreateBill(userId, "abc", 22)
		if err != nil {
			t.Fatal(err)
		}

		b2, err := svc.CreateBill(userId, "def", 13)
		if err != nil {
			t.Fatal(err)
		}

		bills, err := svc.GetBills(userId)
		if err != nil {
			t.Fatal(err)
		}

		if len(bills) != 2 {
			t.Fatal("len(bills) != 2")
		}

		checkBillIsEqual(t, b1, &bills[0])
		checkBillIsEqual(t, b2, &bills[1])
	})
}

func TestBillingService_GetPayment(t *testing.T) {
	t.Run("should return payment", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		userId := "1234"
		amount := 10

		p1, err := svc.CreatePayment(userId, amount)
		if err != nil {
			t.Fatal(err)
		}

		payment, err := svc.GetPayment(p1.ID)
		if err != nil {
			t.Fatal(err)
		}

		checkPaymentIsEqual(t, p1, payment)
	})

	t.Run("should return nil when not found", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		p, err := svc.GetPayment(1)
		if err != nil {
			t.Fatal(err)
		}

		if p != nil {
			t.Fatal("p != nil")
		}
	})
}

func TestBillingService_GetPayments(t *testing.T) {
	t.Run("should return payments", func(t *testing.T) {
		svc := setupService(t)
		defer cleanupService()

		userId := "1234"

		p1, err := svc.CreatePayment(userId, 10)
		if err != nil {
			t.Fatal(err)
		}

		p2, err := svc.CreatePayment(userId, 10)
		if err != nil {
			t.Fatal(err)
		}

		payments, err := svc.GetPayments(userId)
		if err != nil {
			t.Fatal(err)
		}

		if len(payments) != 2 {
			t.Fatal("len(payments) != 2")
		}

		checkPaymentIsEqual(t, p1, &payments[0])
		checkPaymentIsEqual(t, p2, &payments[1])
	})
}
